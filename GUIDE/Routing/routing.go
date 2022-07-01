package Routing

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutesRouting(app *fiber.App) {
	// # Paths
		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("root")
		})
		app.Get("/about", func(c *fiber.Ctx) error {
			return c.SendString("about")
		})
		app.Get("/random.txt", func(c *fiber.Ctx) error {
			return c.SendString("random.txt")
		})

	// # Parameters
		// http://localhost:3000/user/dede/books/tes
		app.Get("/user/:name/books/:title", func(c *fiber.Ctx) error {
			fmt.Println("USER / NAME / BOOKS / TITLE")
			fmt.Fprintf(c, "%s\n", c.Params("name"))
			fmt.Fprintf(c, "%s\n", c.Params("title"))
			return nil
		})
		// http://localhost:3000/user/123
		app.Get("/user/+", func(c *fiber.Ctx) error {
			fmt.Println("USER / +")
			return c.SendString(c.Params("+"))
		})
		// http://localhost:3000/user
		app.Get("/dataname/:name?", func(c *fiber.Ctx) error {
			fmt.Println("DATANAME / NAME ?")
			return c.SendString(c.Params("name"))
		})
		// http://localhost:3000/user
		app.Get("/data/*", func(c *fiber.Ctx) error {
			fmt.Println("DATA / *")
			return c.SendString(c.Params("*"))
		})
		// http://localhost:3000/v1/some/resource/name:customVerb
		app.Get("/v1/some/resource/name\\:customVerb", func(c *fiber.Ctx) error {
			fmt.Println("V1 / SOME / RESOURCE / NAME : customVerb")
			return c.SendString("Hello, Community")
		})
		// http://localhost:3000/plantae/dede.syarifudin
		app.Get("/plantae/:genus.:species", func(c *fiber.Ctx) error {
			fmt.Println("PLANTAE / dede.syarifudin")
			fmt.Fprintf(c, "%s.%s\n", c.Params("genus"), c.Params("species"))
			return nil // prunus.persica
		})
		// http://localhost:3000/flights/sby-jkt
		app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
			fmt.Println("FLIGHTS / sby-jkt")
			fmt.Fprintf(c, "%s-%s\n", c.Params("from"), c.Params("to"))
			return nil // sby-jkt
		})
		// http://localhost:3000/shop/product/color:blue/size:xs
		app.Get("/shop/product/color::color/size::size", func(c *fiber.Ctx) error {
			fmt.Fprintf(c, "%s:%s\n", c.Params("color"), c.Params("size"))
			return nil // blue:xs
		})
		// http://localhost:3000/api-v1
		app.Get("/api-:name", func(c *fiber.Ctx) error {
			fmt.Println("/ API-:NAME")
			return nil
		})
		// http://localhost:3000/@v1
		app.Get("/:sign:param", func(c *fiber.Ctx) error {
			fmt.Println("/ :SIGN :PARAM")
			return nil
		})
		// http://localhost:3000/customer/v1/cart/proxy
		app.Get("/*v1*/proxy", func(c *fiber.Ctx) error {
			fmt.Println("/* V1 */ PROXY")
			return nil
		})
		// http://localhost:3000/v1/brand/4/shop/blue/xs
		app.Get("/v1/*/shop/*", func(c *fiber.Ctx) error {
			fmt.Println("/ V1 / * / SHOP / *")
			return nil
		})

	// # Middleware

	// # Grouping

}

