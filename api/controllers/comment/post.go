package comment

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ArtikelIdParams struct {
    ArtikelId string `query:"artikelId"`    
}

func PostComment (c *fiber.Ctx) error {

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}

    var artikelId ArtikelIdParams
    
    if err := c.QueryParser(&artikelId); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "Invalid query parameters",
		})
    }

    var artikel models.Artikel

    if err := initialize.DB.Where("id = ?", artikelId).First(&artikel).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data artikel cannot be found.",
		})
	}

    thumbnail := form.File["thumbnail"]

    comment := models.Comment {
        Text: form.Value["text"][0],
        Email: form.Value["email"][0],        		
        ArtikelID: artikel.ID,
    }    

	if uuidStr, err := uuid.NewUUID(); err == nil {
		comment.ID = uuidStr
	} else {
		return err
	}

    validate := validator.New()
	
    if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(&comment); err != nil {
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

	 if len(thumbnail) > 0 {			
		uploadedFileNames, err := utils.UploadFile(thumbnail, initialize.ENV_DIR_COMMENT_FILES, comment.ID.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
				"message": "Internal server error.",
			})
		}
		comment.Image = &uploadedFileNames
	}
	
	if result := initialize.DB.Create(&comment); result.Error != nil {		
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to create data comment",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "id": comment.ID,
        "status": "success",        
        "message": "Data comment created.",
    })
}