package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type PostRequestHandlerWithExtra[T any, Q any] func(body T, extra Q) (any, error)
type PostRequestHandler[T any] func(body T) (any, error)
type GetRequestHandlerWithExtra[Q any] func(extra Q) (any, error)
type GetRequestHandler func() (any, error)
type GetExtra[Q any] func(ctx *fiber.Ctx) (Q, error)

// A Simple Post Request with a typed param
func Post[T any](group *ApiGroup, routeName string, body T, handler PostRequestHandler[T]) {
	wrappedHandler := func(body T, extra string) (any, error) {
		return handler(body)
	}
	extraFunc := func(ctx *fiber.Ctx) (string, error) {
		return "", nil
	}
	PostWithExtra(group, routeName, body, wrappedHandler, extraFunc)
}

// A Post Request with a typed param body and an extra function using the Context
func PostWithExtra[T any, Q any](group *ApiGroup, routeName string, body T, handler PostRequestHandlerWithExtra[T, Q], extraFunc GetExtra[Q]) {
	log.Println("Registering route", routeName)

	// Blindly append the route to the array of routes for the group
	group.Routes = append(group.Routes, routeInfo{
		RouteName: routeName,
		Body:      body,
		RouteType: "POST",
		GroupName: group.Name,
	})

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
}

func Get(group *ApiGroup, routeName string, handler GetRequestHandler) {
	wrappedHandler := func(extra string) (any, error) {
		return handler()
	}
	extraFunc := func(ctx *fiber.Ctx) (string, error) {
		return "", nil
	}
	GetWithExtra(group, routeName, wrappedHandler, extraFunc)
}

func GetWithExtra[Q any](group *ApiGroup, routeName string, handler GetRequestHandlerWithExtra[Q], extraFunc GetExtra[Q]) {
	log.Println("Registering route", routeName)
	group.Routes = append(group.Routes, routeInfo{
		RouteName: routeName,
		RouteType: "GET",
		GroupName: group.Name,
	})

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
}
