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

type PostRequestHandlerWithExtra[T any, R any, Q any] func(body T, extra Q) (*R, error)
type PostRequestHandler[T any, R any] func(body T) (*R, error)
type GetRequestHandlerWithExtra[R any, Q any] func(extra Q) (*R, error)
type GetRequestHandler[R any] func() (*R, error)
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

func (r *RouteInfo) WithReturnType(returnValue any) *RouteInfo {
	r.Returns = fiber.Map{
		"status_code": 200,
		"data":        returnValue,
	}
	return r
}

func (r *RouteInfo) WithBodyType(bodyType any) *RouteInfo {
	r.Body = bodyType
	return r
}

// A Simple Post Request with a typed param
func Post[T any, R any](group *ApiGroup, routeName string, handler PostRequestHandler[T, R]) *RouteInfo {
	wrappedHandler := func(body T, extra string) (*R, error) {
		return handler(body)
	}
	extraFunc := func(ctx *fiber.Ctx) (string, error) {
		return "", nil
	}
	return PostWithExtra(group, routeName, wrappedHandler, extraFunc)
}

// A Post Request with a typed param body and an extra function using the Context
func PostWithExtra[T, R, Q any](group *ApiGroup, routeName string, handler PostRequestHandlerWithExtra[T, R, Q], extraFunc GetExtra[Q]) *RouteInfo {
	log.Println("Registering route", group.Name, routeName)
	// Blindly append the route to the array of routes for the group
	routeInfo := RouteInfo{
		RouteName: routeName,
		Body:      new(T),
		RouteType: "POST",
		GroupName: group.Name,
		HasQuery:  false,
		HasParams: false,
		Queries:   []string{},
		Params:    []string{},
		Returns: fiber.Map{
			"status_code": 200,
			"data":        new(R),
		},
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

func Get[R any](group *ApiGroup, routeName string, handler GetRequestHandler[R]) *RouteInfo {
	wrappedHandler := func(extra string) (*R, error) {
		return handler()
	}
	extraFunc := func(ctx *fiber.Ctx) (string, error) {
		return "", nil
	}
	return GetWithExtra(group, routeName, wrappedHandler, extraFunc)
}

func GetWithExtra[R, Q any](group *ApiGroup, routeName string, handler GetRequestHandlerWithExtra[R, Q], extraFunc GetExtra[Q]) *RouteInfo {
	log.Println("Registering route", routeName)
	routeInfo := RouteInfo{
		RouteName: routeName,
		RouteType: "GET",
		GroupName: group.Name,
		Returns: fiber.Map{
			"status_code": 200,
			"data":        new(R),
		},
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

type RawRequestHandler = func(*fiber.Ctx) error

func RawGet(group *ApiGroup, routeName string, handler RawRequestHandler) *RouteInfo {
	log.Println("Registering route", routeName)
	routeInfo := RouteInfo{
		RouteName: routeName,
		RouteType: "GET",
		GroupName: group.Name,
		Returns:   "Unknown",
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
		return handler(ctx)
	})

	return &routeInfo
}

func RawPost(group *ApiGroup, routeName string, handler RawRequestHandler) *RouteInfo {
	log.Println("Registering route", routeName)
	routeInfo := RouteInfo{
		RouteName: routeName,
		RouteType: "POST",
		GroupName: group.Name,
		Returns:   "Unknown",
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

	group.Ctx.Post(routeName, func(ctx *fiber.Ctx) error {
		defer handlePanic(ctx)
		return handler(ctx)
	})

	return &routeInfo
}
