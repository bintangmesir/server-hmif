package artikel

import (
	"server/initialize"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetByIdArtikel (c *fiber.Ctx) error {

    id := c.Params("id")

    var artikel models.Artikel

    if err := initialize.DB.Where("id = ?", id).Preload("Admins", "ArtikelContents", "Comment").First(&artikel).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data artikel cannot be found.",
		})
	}
    
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    artikel,
    })
}