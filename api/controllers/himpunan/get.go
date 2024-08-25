package himpunan

import (
	"server/initialize"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
)


func GetHimpunan(c *fiber.Ctx) error {
    
    var himpunan models.Himpunan
    
    if error := initialize.DB.Model(&models.Himpunan{}).First(&himpunan).Error; error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to retrieve admins",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "data": himpunan,
    })
}