package artikelcontent

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteArtikelContent (c *fiber.Ctx) error {

    id := c.Params("id")

    var artikelContent models.ArtikelContent

    if err := initialize.DB.Where("id = ?", id).First(&artikelContent).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data artikel content cannot be found.",
		})
	}

    if artikelContent.Tipe == "image" {        
        if  err := utils.DeleteFile(&artikelContent.Content, initialize.ENV_DIR_ARTIKEL_FILES, id); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
                "message": "Internal server error.",
            })
        }        
    }

    if result := initialize.DB.Delete(&artikelContent); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to delete data artikel content",
        })
	}

    
    return c.SendStatus(fiber.StatusNoContent)
}