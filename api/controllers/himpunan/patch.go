package himpunan

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PatchHimpunanBody struct {
    JumlahPengurus    int64     `gorm:"not null" validate:"required" json:"jumlahPengurus"`
    JumlahMahasiswa   int64     `gorm:"not null" validate:"required" json:"jumlahMahasiswa"`
    JumlahDepartemen  int64     `gorm:"not null" validate:"required" json:"jumlahDepartemen"`
    NamaProker        string    `gorm:"size:255;not null" json:"namaProker"`    
}

func PatchHimpunan (c *fiber.Ctx) error {

    id := c.Params("id")

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	var himpunan models.Himpunan
    
	fotoProfile := form.File["galeriMahasiswa"]

    jumlahPengurus, err := strconv.Atoi(form.Value["jumlahPengurus"][0]); if err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to convert variable jumlah"})
    }

    jumlahMahasiswa, err := strconv.Atoi(form.Value["jumlahMahasiswa"][0]); if err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to convert variable jumlah"})
    }

    JumlahDepartemen, err := strconv.Atoi(form.Value["jumlahDepartemen"][0]); if err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to convert variable jumlah"})
    }
        
    newHimpunan := PatchHimpunanBody {
        JumlahPengurus: int64(jumlahPengurus),
        JumlahMahasiswa: int64(jumlahMahasiswa),  
        JumlahDepartemen: int64(JumlahDepartemen),         
        NamaProker: form.Value["namaProker"][0],
    }    

    if err := initialize.DB.Where("id = ?", id).First(&himpunan).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data Himpunan cannot be found.",
		})
	}

    validate := validator.New()
	
    if err := c.BodyParser(&newHimpunan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}	

	if err := validate.Struct(&newHimpunan); err != nil {
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
	
	himpunan = models.Himpunan{
		JumlahPengurus: newHimpunan.JumlahPengurus,
		JumlahMahasiswa: newHimpunan.JumlahMahasiswa,
		JumlahDepartemen: newHimpunan.JumlahDepartemen,
		NamaProker: newHimpunan.NamaProker,
	}

	if len(fotoProfile) > 0 {		        
        if himpunan.GaleriMahasiswa != nil {
            if  err := utils.DeleteFile(himpunan.GaleriMahasiswa, initialize.ENV_DIR_HIMPUNAN_FILES, id); err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                    "status": "error",
                    "message": "Internal server error.",
                })
		    }

        }
		uploadedFileNames, err := utils.UploadFile(fotoProfile, initialize.ENV_DIR_HIMPUNAN_FILES, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
				"message": "Internal server error.",
			})
		}
		himpunan.GaleriMahasiswa = &uploadedFileNames
	}

	if result := initialize.DB.Where("id = ?", id).Updates(&himpunan); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to update data himpunan",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Data himpunan updated.",
    })
}