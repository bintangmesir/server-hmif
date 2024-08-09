package barang

import (
	"server/initialize"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
)


func DeleteBarang (c *fiber.Ctx) error {

    id := c.Params("id")

    var barang models.Barang

    if err := initialize.DB.Where("id = ?", id).First(&barang).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data barang cannot be found.",
		})
	}

    if result := initialize.DB.Delete(&barang); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to delete data barang",
        })
	}

    
    return c.SendStatus(fiber.StatusNoContent)
}