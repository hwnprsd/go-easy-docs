package fiberw

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func handlePanic(c *fiber.Ctx) {
	// Panic handler
	if err := recover(); err != nil {
		fmt.Println("We survived a panic")
		fmt.Println(err)
		c.Status(500).JSON(fiber.Map{
			"status_code": 500,
			"message":     "A server panic has occured",
			"error":       err,
		})
	}
}

type PostRequestHandlerWithExtra[T any, Q any] func(body T, extra Q) (any, error)
type PostRequestHandler[T any] func(body T) (any, error)
type GetRequestHandlerWithExtra[Q any] func(extra Q) (any, error)
type GetRequestHandler func() (any, error)
type GetExtra[Q any] func(ctx *fiber.Ctx) (Q, error)

func (r *RouteInfo) WithQuery(query string) *RouteInfo {
	r.HasQuery = true
	r.Queries = append(r.Queries, query)
	return r
}
func (r *RouteInfo) WithParam(param string) *RouteInfo {
	r.HasParams = true
	r.Params = append(r.Queries, param)
	return r
}

// A Simple Post Request with a typed param
func Post[T any](group *ApiGroup, routeName string, body T, handler PostRequestHandler[T]) *RouteInfo {
	wrappedHandler := func(body T, extra string) (any, error) {
		return handler(body)
	}
	extraFunc := func(ctx *fiber.Ctx) (string, error) {
		return "", nil
	}
	return PostWithExtra(group, routeName, body, wrappedHandler, extraFunc)
}

// A Post Request with a typed param body and an extra function using the Context
func PostWithExtra[T any, Q any](group *ApiGroup, routeName string, body T, handler PostRequestHandlerWithExtra[T, Q], extraFunc GetExtra[Q]) *RouteInfo {
	log.Println("Registering route", routeName)

	// Blindly append the route to the array of routes for the group
	routeInfo := RouteInfo{
		RouteName: routeName,
		Body:      body,
		RouteType: "POST",
		GroupName: group.Name,
		HasQuery:  false,
		HasParams: false,
		Queries:   []string{},
		Params:    []string{},
	}
	group.Routes = append(group.Routes, &routeInfo)

	exists, index := false, -1
	for i, g := range ApplicationData.Groups {
		if g.Name == group.Name {
			exists = true
			index = i
		}
	}
	if !exists {
		ApplicationData.Groups = append(ApplicationData.Groups, *group)
	} else {
		ApplicationData.Groups[index] = *group
	}
	// Validate Body

	group.Ctx.Post(routeName, func(ctx *fiber.Ctx) error {
		defer handlePanic(ctx)
		data := new(T)
		log.Println("Before parsing")
		log.Println(data)
		err := ctx.BodyParser(data)
		if err != nil {
			panic(err)
		}
		log.Println("After parsing")
		log.Println(data)
		if b, err := ValidateBody(data, ctx); b {
			return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"status_code": 400,
				"error":       err,
			})
		}
		extra, err := extraFunc(ctx)
		if err != nil {
			log.Printf("Err")
			return handleError(err, ctx)
		}
		response, err := handler(*data, extra)
		if err != nil {
			return handleError(err, ctx)
		}
		ctx.Status(200).JSON(fiber.Map{
			"status_code": 200,
			"data":        response,
		})
		return nil
	})
	return &routeInfo
}

func Get(group *ApiGroup, routeName string, handler GetRequestHandler) *RouteInfo {
	wrappedHandler := func(extra string) (any, error) {
		return handler()
	}
	extraFunc := func(ctx *fiber.Ctx) (string, error) {
		return "", nil
	}
	return GetWithExtra(group, routeName, wrappedHandler, extraFunc)
}

func GetWithExtra[Q any](group *ApiGroup, routeName string, handler GetRequestHandlerWithExtra[Q], extraFunc GetExtra[Q]) *RouteInfo {
	log.Println("Registering route", routeName)
	routeInfo := RouteInfo{
		RouteName: routeName,
		RouteType: "GET",
		GroupName: group.Name,
	}
	group.Routes = append(group.Routes, &routeInfo)

	exists, index := false, -1
	for i, g := range ApplicationData.Groups {
		if g.Name == group.Name {
			exists = true
			index = i
		}
	}
	if !exists {
		ApplicationData.Groups = append(ApplicationData.Groups, *group)
	} else {
		ApplicationData.Groups[index] = *group
	}

	group.Ctx.Get(routeName, func(ctx *fiber.Ctx) error {
		defer handlePanic(ctx)
		extra, err := extraFunc(ctx)
		if err != nil {
			log.Printf("Error handling extrafunc")
			return handleError(err, ctx)
		}
		response, err := handler(extra)
		if err != nil {
			return handleError(err, ctx)
		}
		ctx.Status(200).JSON(fiber.Map{
			"status_code": 200,
			"data":        response,
		})
		return nil
	})
	return &routeInfo
}
