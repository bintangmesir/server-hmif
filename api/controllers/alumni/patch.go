package alumni

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PatchAlumniBody struct {
	Angkatan        string    `gorm:"size:4" json:"angkatan" validate:"required,max=4"`
    Nama            string    `gorm:"size:100" json:"nama" validate:"required,max=100"`
    NoTelephone     string    `gorm:"size:15" json:"noTelephone" validate:"required,max=15"`
}

func PatchAlumni (c *fiber.Ctx) error {

    id := c.Params("id")

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}    

    var alumni models.Alumni
        
    newAlumni := PatchAlumniBody {
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

	if err := validate.Struct(&newAlumni); err != nil {
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
