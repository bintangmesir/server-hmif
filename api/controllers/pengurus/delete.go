package pengurus

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func DeletePengurus (c *fiber.Ctx) error {

    id := c.Params("id")

    var pengurus models.Pengurus

    if err := initialize.DB.Where("id = ?", id).First(&pengurus).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data pengurus cannot be found.",
		})
	}

     if pengurus.Foto != nil {
        if  err := utils.DeleteFile(pengurus.Foto, initialize.ENV_DIR_PENGURUS_FILES, id); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
                "message": "Internal server error.",
            })
		}
    }

    if result := initialize.DB.Delete(&pengurus); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to delete data pengurus",
        })
	}

    
    return c.SendStatus(fiber.StatusNoContent)
}