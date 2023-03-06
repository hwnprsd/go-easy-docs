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

type CreateUserDto struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     uint   `json:"age"`
}

type User struct {
	ID      string `json:"ID"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     uint   `json:"age"`
}

func (s SwaggerApp) testHandler(body TestBody) (any, error) {
	log.Println("Body", body.Foo)
	return body, nil
}

func (s SwaggerApp) newTestHandler(body AddDaoDto) (any, error) {
	return "OKAY", nil
}

type CraeteDDQuery struct {
	Address  string
	Quantity uint
}

func (s SwaggerApp) CreateDataDao(body AddDaoDto, user User) (*interface{}, error) {
	return nil, nil
}

func (SwaggerApp) getTest() (any, error) {
	return "Looks Good", nil
}

func CreateUser(userData *CreateUserDto) (*User, error) {
	return &User{
		Name:    "Roger",
		Address: "21, Palm Drive",
		Age:     80,
	}, nil
}

func CreateUserGet() (*User, error) {
	return &User{
		Name:    "Roger",
		Address: "21, Palm Drive",
		Age:     80,
	}, nil
}

func AuthMiddleware(ctx *fiber.Ctx) (User, error) {
	// Check Auth
	return User{}, nil
}

func Run() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	sApp := SwaggerApp{}

	testGroup := fiberw.NewGroup(app, "/test")

	fiberw.PostWithExtra(testGroup, "/pathMe2/:address", sApp.CreateDataDao, AuthMiddleware).WithParam("address").WithQuery("maze").WithParam("rand")

	getGroup := fiberw.NewGroup(app, "/users")

	fiberw.Post(getGroup, "/create-user", CreateUser)
	fiberw.Get(getGroup, "/create-user", CreateUserGet)
	fiberw.Post(getGroup, "/create-new-user", CreateUser)
	// routes := app.GetRoutes()

	app.Get("/docs", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		docs := fiberw.GenerateDocs("Test Application", "Used to test go-auto-docs")
		c.Render("index", docs)
		// fiberw.WriteApiDocumentation("Test Application", "Used to test go-auto-swagger", c)
		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
