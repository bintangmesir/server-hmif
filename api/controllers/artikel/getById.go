package artikel

import (
	"server/initialize"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetByIdArtikel (c *fiber.Ctx) error {

    id := c.Params("id")

	var newArtikel models.Artikel

	if err := initialize.DB.Where("id = ?", id).First(&newArtikel).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data artikel cannot be found.",
		})
	}

	newArtikel.View = int64(newArtikel.View + 1)

    if result := initialize.DB.Where("id = ?", id).Updates(&newArtikel); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to update data artikel",
        })
	}

    var artikel models.Artikel

    if err := initialize.DB.Where("id = ?", id).Preload("Admins").Preload("ArtikelContents").Preload("Comment").First(&artikel).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data artikel cannot be found.",
		})
	}
    
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    artikel,
    })
}