package youtube

import (
	"server/initialize"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteYoutube (c *fiber.Ctx) error {

    id := c.Params("id")

    var youtube models.Youtube

    if err := initialize.DB.Where("id = ?", id).First(&youtube).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data youtube cannot be found.",
		})
	}

    if result := initialize.DB.Delete(&youtube); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to delete data youtube",
        })
	}
    
    return c.SendStatus(fiber.StatusNoContent)
}