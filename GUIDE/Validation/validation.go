package Validation

import (
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutesValidation(app *fiber.App) {

}