package alumni

import (
	"server/initialize"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteAlumni (c *fiber.Ctx) error {

    id := c.Params("id")

    var alumni models.Alumni

    if err := initialize.DB.Where("id = ?", id).First(&alumni).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data alumni cannot be found.",
		})
	}

    if result := initialize.DB.Delete(&alumni); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to delete data alumni",
        })
	}

    
    return c.SendStatus(fiber.StatusNoContent)
}