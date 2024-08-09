package buku

import (
	"server/initialize"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetByIdBuku (c *fiber.Ctx) error {

    id := c.Params("id")

    var buku models.Buku

    if err := initialize.DB.Where("id = ?", id).First(&buku).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data buku cannot be found.",
		})
	}
    
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    buku,
    })
}