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
	Foo      string `json:"foo" validate:"required"`
	FancyFoo string `json:"fancy_foo_x" validate:"required"`
}

type AddDaoDto struct {
	DaoName string `json:"dao_name"`
	DaoId   uint   `json:"dao_id"`
}

func (s SwaggerApp) testHandler(body TestBody) (any, error) {
	log.Println("Body", body.Foo)
	return body, nil
}

func (s SwaggerApp) newTestHandler(body AddDaoDto) (any, error) {
	return "OKAY", nil
}

func (SwaggerApp) getTest() (any, error) {
	return "Looks Good", nil
}

func Run() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	sApp := SwaggerApp{}

	testGroup := utils.NewGroup(app, "/test")
	utils.Post(testGroup, "/submit", TestBody{}, sApp.testHandler)
	utils.Post(testGroup, "/rekt", TestBody{}, sApp.testHandler)

	getGroup := utils.NewGroup(app, "/users")
	utils.Get(getGroup, "/all", sApp.getTest)

	daoGroup := utils.NewGroup(app, "/daos")
	utils.Post(daoGroup, "/add", AddDaoDto{}, sApp.newTestHandler)

	// routes := app.GetRoutes()

	viewData := utils.GenerateDocs("Runtime Docs in Go", "API Docs generated at runtime using HTML Templates and a very simple data structure")

	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Render("index", viewData)
	})

	log.Fatal(app.Listen(":3000"))
}
