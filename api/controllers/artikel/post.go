package artikel

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostArtikel (c *fiber.Ctx) error {

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}

    thumbnail := form.File["thumbnail"]

    commentEnabled, err := strconv.ParseBool(form.Value["commentEnabled"][0]); if err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to convert Comment Enabled variable "})
    }

    artikel := models.Artikel {
        Title: form.Value["title"][0],
        SubTitle: form.Value["subTitle"][0],
        CommentEnabled: commentEnabled,  
		View: 1,      	
    }    

	if uuidStr, err := uuid.NewUUID(); err == nil {
		artikel.ID = uuidStr
	} else {
		return err
	}

    validate := validator.New()
	
    if err := c.BodyParser(&artikel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(artikel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": err.Error(),
        })
	}

	 if len(thumbnail) > 0 {			
		uploadedFileNames, err := utils.UploadFile(thumbnail, initialize.ENV_DIR_ARTIKEL_FILES, artikel.ID.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
				"message": "Internal server error.",
			})
		}
		artikel.Thumbnail = &uploadedFileNames
	}
	
	if result := initialize.DB.Create(&artikel); result.Error != nil {		
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to create data artikel",
        })
	}


    refreshToken := c.Cookies("refreshToken")

    claims, err := utils.ParseToken(refreshToken, "refresh")
    if err != nil || claims["type"] != "refresh" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "status": "error",
            "message": "You are not authenticated.",
        })
    }
	
    userID := claims["id"]

	var admin models.Admin

	initialize.DB.Model(&artikel).Association("Admins").Append([]models.Admin{admin})

	if err := initialize.DB.Where("id = ?", userID).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data admin cannot be found.",
		})
	}
	

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "id": artikel.ID,
        "status": "success",        
        "message": "Data artikel created.",
    })
}