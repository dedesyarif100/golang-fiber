package ErrorHandling

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutesErrorHandling(app *fiber.App) {
	// # Catching Errors
		app.Use(recover.New())
		app.Get("/", func(c *fiber.Ctx) error {
			// panic("This panic is caught by fiber")
			// return c.SendFile("file-does-not-exist")
			// return fiber.ErrServiceUnavailable
			return fiber.NewError(fiber.StatusServiceUnavailable, "On vacation!")
		})

	// # Default Error Handler
		// Default error handler
		// var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
		// 	// Default 500 statuscode
		// 	code := fiber.StatusInternalServerError

		// 	if e, ok := err.(*fiber.Error); ok {
		// 		// Override status code if fiber.Error type
		// 		code = e.Code
		// 	}
		// 	// Set Content-Type: text/plain; charset=utf-8
		// 	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

		// 	// Return statuscode with error message
		// 	return c.Status(code).SendString(err.Error())
		// }
		// app.Get("default-error", DefaultErrorHandler)

		apps := fiber.New(fiber.Config{
			// Override default error handler
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				// Status code defaults to 500
				code := fiber.StatusInternalServerError
		
				// Retrieve the custom status code if it's an fiber.*Error
				if e, ok := err.(*fiber.Error); ok {
					code = e.Code
				}
		
				// Send custom error page
				err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
				if err != nil {
					// In case the SendFile fails
					return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
				}
		
				// Return from handler
				return nil
			},
		})
		apps.Get("/apps", func(c *fiber.Ctx) error {
			fmt.Println("APPS")
			return nil
		})


}


