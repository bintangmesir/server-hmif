package youtube

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PatchYoutubeBody struct {
	Judul       string      `gorm:"size:100" json:"judul" validate:"required,max=100"`
    Link        string      `gorm:"size:255" json:"link" validate:"required,url,max=100"`
}

func PatchYoutube (c *fiber.Ctx) error {

    id := c.Params("id")

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}    

    var youtube models.Youtube
        
    newYoutube := PatchYoutubeBody {
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
		var errorMassage []string

		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors{			
			errorMassage = append(errorMassage, utils.ErrorMassage(fieldError.Field(), fieldError.Tag(), fieldError.Param()))
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": errorMassage,
        })
	}
	
    youtube = models.Youtube{
        Judul: newYoutube.Judul,
        Link: newYoutube.Link,
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