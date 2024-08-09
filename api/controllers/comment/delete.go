package comment

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/gofiber/fiber/v2"
)


func DeleteComment (c *fiber.Ctx) error {

    id := c.Params("id")

    var comment models.Comment

    if err := initialize.DB.Where("id = ?", id).First(&comment).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data comment cannot be found.",
		})
	}

    if comment.Image != nil {
        if  err := utils.DeleteFile(comment.Image, initialize.ENV_DIR_COMMENT_FILES, id); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
                "message": "Internal server error.",
            })
		}
    }

    if result := initialize.DB.Delete(&comment); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to delete data comment",
        })
	}

    
    return c.SendStatus(fiber.StatusNoContent)
}