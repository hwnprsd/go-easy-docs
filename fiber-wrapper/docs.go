package fiberw

import (
	"github.com/gofiber/fiber/v2"
)

// API Group with which the docs will be generated
type ApiGroup struct {
	Name   string
	Routes []routeInfo
	Ctx    fiber.Router
}

// New API Group
func NewGroup(app *fiber.App, name string) *ApiGroup {
	return &ApiGroup{Name: name, Routes: []routeInfo{}, Ctx: app.Group(name)}
}

// Api Doc Config
type ApiDoc struct {
	ApplicationName string
	Description     string
	BaseUrl         string
	Groups          []ApiGroup
}

// Info a route carries
type routeInfo struct {
	RouteName   string
	Body        any
	Headers     any
	RouteType   string
	Description string
	GroupName   string
}

// Application data of the whole server
// All the data is stored in this public variable. Please use PostRequest and GetRequest wrappers provided by this
// library for the API Docs to be generated
var ApplicationData = ApiDoc{Groups: []ApiGroup{}}

func GenerateDocs(applicationName string, description string) interface{} {
	ApplicationData.ApplicationName = applicationName
	ApplicationData.Description = description
	return ApplicationData
}
