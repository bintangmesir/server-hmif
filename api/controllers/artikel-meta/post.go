package artikelmeta

import (
	"server/initialize"
	"server/internal/models"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ArtikelIdParams struct {
    ArtikelId string `query:"artikelId"`    
}

func PostArtikelLike (c *fiber.Ctx) error {

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

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}

    like, err := strconv.Atoi(form.Value["like"][0]); if err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to convert variable like",
        })
    }

    artikelMeta := models.ArtikelMeta {
        Like: int64(like),
        Email: form.Value["email"][0],        
    }    

	if uuidStr, err := uuid.NewUUID(); err == nil {
		artikelMeta.ID = uuidStr
	} else {
		return err
	}    
    
    
    validate := validator.New()
	
    if err := c.BodyParser(&artikelMeta); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(&artikelMeta); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": err.Error(),
        })
	}
	    
	if result := initialize.DB.Create(&artikelMeta); result.Error != nil {		
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to create data artikel meta",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{        
        "status": "success",        
    })
}