package pengurus

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PatchPengurusBody struct {
	Name        string      `gorm:"size:100;not null" json:"name" validate:"required,max=100"`
    Departemen  string      `gorm:"size:20;not null" json:"departemen" validate:"required,oneof=kahim_wakahim sekretaris bendahara departemen_iptek departemen_kominfo departemen_kaderisasi departemen_prhp departemen_pengmas"`
    Jabatan     string      `gorm:"size:20;not null" json:"jabatan" validate:"required,oneof=ketua_himpunan wakil_ketua_himpunan sekretaris_1 sekretaris_2 bendahara_1 bendahara_2 kepala_departemen staff_departemen"`    
}


func PatchPengurus (c *fiber.Ctx) error {

    id := c.Params("id")

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}

    foto := form.File["foto"]

    var pengurus models.Pengurus

    newPengurus := PatchPengurusBody {
        Name: form.Value["name"][0],
        Departemen: form.Value["departemen"][0],
        Jabatan: form.Value["jabatan"][0],   
    }    

    if err := initialize.DB.Where("id = ?", id).First(&pengurus).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data pengurus cannot be found.",
		})
	}

    validate := validator.New()
	
    if err := c.BodyParser(&newPengurus); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(&newPengurus); err != nil {
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

    pengurus = models.Pengurus{
        Name: newPengurus.Name,
        Departemen: newPengurus.Departemen,
        Jabatan: newPengurus.Jabatan,   
    }
        
    if len(foto) > 0 {		        
        if pengurus.Foto != nil {
            if  err := utils.DeleteFile(pengurus.Foto, initialize.ENV_DIR_PENGURUS_FILES, id); err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                    "status": "error",
                    "message": "Internal server error.",
                })
		    }

        }
		uploadedFileNames, err := utils.UploadFile(foto, initialize.ENV_DIR_PENGURUS_FILES, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
				"message": "Internal server error.",
			})
		}
		pengurus.Foto = &uploadedFileNames
	}
	
	if result := initialize.DB.Where("id = ?", id).Updates(&pengurus); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to update data pengurus",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Data pengurus updated.",
    })
}