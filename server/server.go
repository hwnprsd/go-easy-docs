package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/hwnprsd/go-api-docs/utils"
)

type SwaggerApp struct {
	App *fiber.App
	Db  any
}

type TestBody struct {
	Foo      string `json:"foo"`
	FancyFoo string `json:"FancyFoo"`
}

func (s SwaggerApp) testHandler(body TestBody) (any, error) {
	log.Println(body.Foo)
	return "OKAY", nil
}

func Run() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	sApp := SwaggerApp{}

	testGroup := utils.NewGroup(app, "/test")
	utils.Post(testGroup, "/submit", TestBody{}, sApp.testHandler, nil)
	utils.Post(testGroup, "/rekt", TestBody{}, sApp.testHandler)

	// routes := app.GetRoutes()

	viewData := utils.GenerateDocs("Go Docs", "This test")
	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Render("index", viewData)
	})
	log.Fatal(app.Listen(":3000"))
}
