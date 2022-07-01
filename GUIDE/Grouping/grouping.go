package Grouping

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutesGrouping(app *fiber.App) {
	// # Paths

	// # Group Handlers
		handler := func(c *fiber.Ctx) error {
			fmt.Println("MASUK GROUPING")
			return c.SendStatus(fiber.StatusOK)
		}
		api := app.Group("/api") // /api
		v1 := api.Group("/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
			fmt.Println("TES")
			c.Set("Version", "v1")
			return c.Next()
		})
		v1.Get("/list", handler) // /api/v1/list
		v1.Get("/user", handler) // /api/v1/user
}

