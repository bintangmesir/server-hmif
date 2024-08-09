package youtube

import (
	"server/initialize"
	"server/internal/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostYoutube (c *fiber.Ctx) error {

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}    

    youtube := models.Youtube{
        Judul: form.Value["judul"][0],
        Link: form.Value["link"][0],    
    }    

    if uuidStr, err := uuid.NewUUID(); err == nil {
		youtube.ID = uuidStr
	} else {
		return err
	}    

    validate := validator.New()
	
    if err := c.BodyParser(&youtube); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(youtube); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": err.Error(),
        })
	}

	if result := initialize.DB.Create(&youtube); result.Error != nil {		
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to create data youtube",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Data youtube created.",
    })
}