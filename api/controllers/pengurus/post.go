package pengurus

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostPengurus (c *fiber.Ctx) error {

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}

    foto := form.File["foto"]

    pengurus := models.Pengurus {
        Name: form.Value["name"][0],
        Departemen: form.Value["departemen"][0],
        Jabatan: form.Value["jabatan"][0],        		
    }    

	if uuidStr, err := uuid.NewUUID(); err == nil {
		pengurus.ID = uuidStr
	} else {
		return err
	}    

    validate := validator.New()
	
    if err := c.BodyParser(&pengurus); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(pengurus); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": err.Error(),
        })
	}

	 if len(foto) > 0 {			
		uploadedFileNames, err := utils.UploadFile(foto, initialize.ENV_DIR_PENGURUS_FILES, pengurus.ID.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
				"message": "Internal server error.",
			})
		}
		pengurus.Foto = &uploadedFileNames
	}
	
	if result := initialize.DB.Create(&pengurus); result.Error != nil {		
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to create data pengurus",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Data pengurus created.",
    })
}