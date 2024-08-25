package admin

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostAdmin (c *fiber.Ctx) error {

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	var admin models.Admin

    fotoProfile := form.File["fotoProfile"]

    newAdmin := models.Admin {
        Name: form.Value["name"][0],
        Email: form.Value["email"][0],
        Password: form.Value["password"][0],        
		Role: form.Value["role"][0],
    }    

	if uuidStr, err := uuid.NewUUID(); err == nil {
		newAdmin.ID = uuidStr
	} else {
		return err
	}

    if err := initialize.DB.Where("email = ?", newAdmin.Email).First(&admin).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
			"message": "Email already used.",
		})
	}

    validate := validator.New()
	
    if err := c.BodyParser(&newAdmin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(&newAdmin); err != nil {
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

	 if len(fotoProfile) > 0 {			
		uploadedFileNames, err := utils.UploadFile(fotoProfile, initialize.ENV_DIR_ADMIN_FILES, newAdmin.ID.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
				"message": "Internal server error.",
			})
		}
		newAdmin.FotoProfile = &uploadedFileNames
	}
	
	if result := initialize.DB.Create(&newAdmin); result.Error != nil {		
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to create data admin",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Data admin created.",
    })
}