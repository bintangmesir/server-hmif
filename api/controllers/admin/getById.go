package admin

import (
	"server/initialize"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetByIdAdmin (c *fiber.Ctx) error {

    id := c.Params("id")

    var admin models.Admin

    if err := initialize.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data admin cannot be found.",
		})
	}
    
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    admin,
    })
}