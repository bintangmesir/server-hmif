package artikel

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteArtikel (c *fiber.Ctx) error {

    id := c.Params("id")

    var artikel models.Artikel

    if err := initialize.DB.Where("id = ?", id).First(&artikel).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data artikel cannot be found.",
		})
	}

     if artikel.Thumbnail != nil {
        if  err := utils.DeleteFile(artikel.Thumbnail, initialize.ENV_DIR_ARTIKEL_FILES, id); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
                "message": "Internal server error.",
            })
		}
    }

    if result := initialize.DB.Delete(&artikel); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to delete data artikel",
        })
	}

    
    return c.SendStatus(fiber.StatusNoContent)
}