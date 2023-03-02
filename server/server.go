package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/hwnprsd/go-easy-docs/fiber-wrapper"
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
	DaoName string `json:"dao_name" validate:"required"`
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

	testGroup := fiberw.NewGroup(app, "/test")
	fiberw.Post(testGroup, "/submit", TestBody{}, sApp.testHandler)
	fiberw.Post(testGroup, "/rekt", TestBody{}, sApp.testHandler)

	getGroup := fiberw.NewGroup(app, "/users")
	fiberw.Get(getGroup, "/all", sApp.getTest)

	daoGroup := fiberw.NewGroup(app, "/daos")
	fiberw.Post(daoGroup, "/add", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add1", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add2", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add3", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add4", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add5", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add6", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add7", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add8", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add9", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add10", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add12", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add13", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add14", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add16", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add17", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add18", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add19", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add20", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add21", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add22", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add23", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add24", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add25", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add26", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add27", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add30", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add31", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add32", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add33", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add34", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add35", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add36", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add37", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add38", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add39", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add41", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add42", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add44", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add45", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add46", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add47", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add48", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add50", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add51", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add52", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add53", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add54", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add55", AddDaoDto{}, sApp.newTestHandler)
	fiberw.Post(daoGroup, "/add56", AddDaoDto{}, sApp.newTestHandler)

	// routes := app.GetRoutes()

	viewData := fiberw.GenerateDocs("Runtime Docs in Go", "API Docs generated at runtime using HTML Templates and a very simple data structure")

	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Render("index", viewData)
	})

	log.Fatal(app.Listen(":3000"))
}
