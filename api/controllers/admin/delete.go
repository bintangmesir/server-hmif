package admin

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteAdmin (c *fiber.Ctx) error {

    id := c.Params("id")

    var admin models.Admin

    if err := initialize.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data admin cannot be found.",
		})
	}

     if admin.FotoProfile != nil {
        if  err := utils.DeleteFile(admin.FotoProfile, initialize.ENV_DIR_ADMIN_FILES, id); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
                "message": "Internal server error.",
            })
		}
    }

    if result := initialize.DB.Delete(&admin); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to delete data admin",
        })
	}

    
    return c.SendStatus(fiber.StatusNoContent)
}