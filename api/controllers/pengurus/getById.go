package pengurus

import (
	"server/initialize"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetByIdPengurus (c *fiber.Ctx) error {

    id := c.Params("id")

    var pengurus models.Pengurus

    if err := initialize.DB.Where("id = ?", id).First(&pengurus).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data pengurus cannot be found.",
		})
	}
    
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    pengurus,
    })
}