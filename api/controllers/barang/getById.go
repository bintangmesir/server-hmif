package barang

import (
	"server/initialize"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetByIdBarang (c *fiber.Ctx) error {

    id := c.Params("id")

    var barang models.Barang

    if err := initialize.DB.Where("id = ?", id).First(&barang).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data barang cannot be found.",
		})
	}
    
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    barang,
    })
}