package alumni

import (
	"server/initialize"
	"server/internal/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostAlumni (c *fiber.Ctx) error {

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}    

    alumni := models.Alumni {
        Angkatan: form.Value["angkatan"][0],
        Nama: form.Value["nama"][0],
        NoTelephone: form.Value["noTelephone"][0],	
    }    

	if uuidStr, err := uuid.NewUUID(); err == nil {
		alumni.ID = uuidStr
	} else {
		return err
	}

    validate := validator.New()
	
    if err := c.BodyParser(&alumni); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(alumni); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": err.Error(),
        })
	}

	if result := initialize.DB.Create(&alumni); result.Error != nil {		
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to create data alumni",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Data alumni created.",
    })
}