package youtube

import (
	"server/initialize"
	"server/internal/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)


func PatchYoutube (c *fiber.Ctx) error {

    id := c.Params("id")

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}    

    var youtube models.Youtube
        
    newYoutube := models.Youtube {
        Judul: form.Value["judul"][0],
        Link: form.Value["link"][0],               
    }    

    if err := initialize.DB.Where("id = ?", id).First(&youtube).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data youtube cannot be found.",
		})
	}

    validate := validator.New()
	
    if err := c.BodyParser(&newYoutube); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(&newYoutube); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": err.Error(),
        })
	}
	
	if result := initialize.DB.Where("id = ?", id).Updates(&youtube); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to update data youtube",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Data youtube updated.",
    })
}