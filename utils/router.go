package utils

import (
	"log"

	"github.com/fatih/structs"
	"github.com/gofiber/fiber/v2"
)

type ApiGroup struct {
	Name   string
	Routes []routeInfo
	Ctx    fiber.Router
}

func NewGroup(app *fiber.App, name string) *ApiGroup {
	return &ApiGroup{Name: name, Routes: []routeInfo{}, Ctx: app.Group(name)}
}

type ApiDoc struct {
	ApplicationName string
	Description     string
	BaseUrl         string
	Groups          []ApiGroup
}

type routeInfo struct {
	RouteName   string
	Body        any
	Headers     any
	RouteType   string
	Description string
	GroupName   string
}

var Data = ApiDoc{Groups: []ApiGroup{}}

type PostRequestHandlerWithCtx[T any] func(body T, ctx fiber.Ctx) (any, error)
type PostRequestHandler[T any] func(body T) (any, error)
type GetRequestHandler[Q any] func(queryParams *Q) (any, error)

func Post[T any](group *ApiGroup, routeName string, body T, handler PostRequestHandler[T]) {
	log.Println("Registering route", routeName)
	group.Routes = append(group.Routes, routeInfo{
		RouteName: routeName,
		Body:      body,
		RouteType: "POST",
		GroupName: group.Name,
	})

	exists, index := false, -1
	for i, g := range Data.Groups {
		if g.Name == group.Name {
			exists = true
			index = i
		}
	}
	if !exists {
		Data.Groups = append(Data.Groups, *group)
	} else {
		Data.Groups[index] = *group
	}
	if queryParams != nil {
		// Parse Query Param
	}
	// Validate Body
	group.Ctx.Post(routeName, func(c *fiber.Ctx) error {
		c.BodyParser(&body)
		response, err := handler(body, ...queryParams)
		if err != nil {
			return err
		}
		c.Status(200).JSON(fiber.Map{
			"status_code": 200,
			"data":        response,
		})
		return nil
	})
}

func Get[Q any](group *ApiGroup, routeName string, handler GetRequestHandler[Q]) {
	log.Println("Registering route", routeName)
}

func GenerateDocs(applicationName string, description string) interface{} {
	Data.ApplicationName = applicationName
	Data.Description = description
	// bytes, _ := json.Marshal(j)
	// log.Println(string(bytes))
	return structs.Map(Data)

}
