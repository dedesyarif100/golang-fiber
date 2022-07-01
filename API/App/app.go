package App

import (
	"fmt"
	"gorm.io/gorm"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) appHandler(c *fiber.Ctx) error {
	fmt.Println("MASUK APP HANDLER")
	return nil
}

func (r *Repository) SetupRoutesApp(app *fiber.App) {
	// # Static
		app.Static("/template", "./template")

	// # Route Handlers
		app.Get("/api/list", func(c *fiber.Ctx) error {
			return c.SendString("I'm a GET request!")
		})
		app.Post("/api/register", func(c *fiber.Ctx) error {
			return c.SendString("I'm a POST request!")
		})
		app.Use("/api", func(c *fiber.Ctx) error {
			fmt.Println("TES")
			return c.Next()
		})

	// # Mount
		micro := fiber.New()
		micro.Get("/doe", func(c *fiber.Ctx) error {
			// fmt.Println("MOUNT")
			return c.SendStatus(fiber.StatusOK)
		})
    	app.Mount("/john", micro) // GET /john/doe -> 200 OK

	// # Group

	// # Route
		app.Route("/test", func(api fiber.Router) {
			api.Get("/foo", r.appHandler).Name("foo") // /test/foo (name: test.foo)
			api.Get("/bar", r.appHandler).Name("bar") // /test/bar (name: test.bar)
		}, "test.")

	// # Server
	
	
	// # HandlersCount
		app.Route("/handleCount", func(api fiber.Router) {
			api.Get("/tes", func(c *fiber.Ctx) error {
				fmt.Println(app.HandlersCount())
				return nil
			})
		})

	// # Stack
		appStack := fiber.New()
		appStack.Get("/john/:age", r.appHandler).Name("john")
		appStack.Post("/register", r.appHandler).Name("register")
		data, _ := json.MarshalIndent(appStack.Stack(), "", "  ")
		app.Get("/stack", func(c *fiber.Ctx) error {
			fmt.Println(string(data))
			return nil
		})

	// # Name
		appName := fiber.New()
		appName.Get("/", r.appHandler)
		appName.Name("index")
		appName.Get("/doe", r.appHandler).Name("home")

		appName.Trace("/tracer", r.appHandler).Name("tracert")
		
		appName.Delete("/delete", r.appHandler).Name("delete")
		
		a := appName.Group("/a")
		a.Name("fd.")
		a.Get("/test", r.appHandler).Name("test")
		
		dataName, _ := json.MarshalIndent(appName.Stack(), "", "  ")
		app.Get("/appName", func(c *fiber.Ctx) error {
			fmt.Print(string(dataName))
			return nil
		})

	// # GetRoute
		appGetRoute := fiber.New()

		appGetRoute.Get("/", r.appHandler).Name("index")

		dataGetRoute, _ := json.MarshalIndent(appGetRoute.GetRoute("index"), "", "  ")
		app.Get("appGetRoute", func(c *fiber.Ctx) error {
			fmt.Print(string(dataGetRoute))
			return nil
		})

	// # Config
		app.Get("appConfig", func(c *fiber.Ctx) error {
			fmt.Println(app.Config())
			return nil
		})

	// # Handler

	// # Listen

	// # ListenTLS

	// # ListenMutualTLS

	// # Listener

	// # Test

	// # Hooks

}

