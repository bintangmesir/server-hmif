package buku

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteBuku (c *fiber.Ctx) error {

    id := c.Params("id")

    var buku models.Buku

    if err := initialize.DB.Where("id = ?", id).First(&buku).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data buku cannot be found.",
		})
	}

     if buku.Cover != nil {
        if  err := utils.DeleteFile(buku.Cover, initialize.ENV_DIR_BUKU_FILES, id); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
                "message": "Internal server error.",
            })
		}
    }

    if result := initialize.DB.Delete(&buku); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to delete data buku",
        })
	}

    
    return c.SendStatus(fiber.StatusNoContent)
}