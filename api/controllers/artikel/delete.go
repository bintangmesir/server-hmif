package artikel

import (
	"fmt"
	"os"
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteArtikel (c *fiber.Ctx) error {

    id := c.Params("id")

    var artikel models.Artikel

    if err := initialize.DB.Where("id = ?", id).Preload("Admins").First(&artikel).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data artikel cannot be found.",
		})
	}

    var admin models.Admin

    if err := initialize.DB.Where("id = ?", artikel.Admins[0].ID).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data admin cannot be found.",
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

    initialize.DB.Model(&artikel).Association("Admins").Delete([]models.Admin{admin})    
    
	var prevArtikelContents []models.ArtikelContent
	if err := initialize.DB.Where("artikel_id = ?", id).Find(&prevArtikelContents).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch existing artikel content.",
		})    
    }

    cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Define the directory path with the given ID
	dirPath := fmt.Sprintf("%s%s%s/", cwd, initialize.ENV_DIR_ARTIKEL_CONTENT_FILES, id)
        
    if result := initialize.DB.Delete(&prevArtikelContents); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to delete data artikel",
        })
	}

    if err := os.RemoveAll(dirPath); err != nil {		
		fmt.Println("Directory not exist...")
	}

    if result := initialize.DB.Delete(&artikel); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to delete data artikel",
        })
	}
        
    return c.SendStatus(fiber.StatusNoContent)
}