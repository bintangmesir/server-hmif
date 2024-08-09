package alumni

import (
	"server/initialize"
	"server/internal/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)


func PatchAlumni (c *fiber.Ctx) error {

    id := c.Params("id")

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}    

    var alumni models.Alumni
        
    newAlumni := models.Alumni {
        Angkatan: form.Value["angkatan"][0],
        Nama: form.Value["nama"][0],                
        NoTelephone: form.Value["noTelephone"][0],                
    }    

    if err := initialize.DB.Where("id = ?", id).First(&alumni).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data alumni cannot be found.",
		})
	}

    validate := validator.New()
	
    if err := c.BodyParser(&newAlumni); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(newAlumni); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": err.Error(),
        })
	}

    alumni = models.Alumni{
        Angkatan: newAlumni.Angkatan,
        Nama: newAlumni.Nama,
        NoTelephone: newAlumni.NoTelephone,
    }
	
	if result := initialize.DB.Where("id = ?", id).Updates(&alumni); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to update data alumni",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Data alumni updated.",
    })
}
