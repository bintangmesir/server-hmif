package artikelcontent

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ArtikelIdParams struct {
    ArtikelId string `query:"artikelId"`    
}

func PostArtikelContent (c *fiber.Ctx) error {

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

     index, err := strconv.Atoi(form.Value["index"][0]); if err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to convert variable index"})
    }

    artikelContent := models.ArtikelContent {
        Index: int64(index),
        Tipe: form.Value["tipe"][0],
        SubTipe: form.Value["subTipe"][0],      
        ArtikelID: artikel.ID,  	
    }    

	if uuidStr, err := uuid.NewUUID(); err == nil {
		artikelContent.ID = uuidStr
	} else {
		return err
	}    
    
    if artikelContent.Tipe == "image" {
        content := form.File["content"]

        if len(content) > 0 {			
            uploadedFileNames, err := utils.UploadFile(content, initialize.ENV_DIR_ARTIKEL_FILES, artikelContent.ID.String())
            if err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                    "status": "error",
                    "message": "Internal server error.",
                })
            }
            artikelContent.Content = uploadedFileNames
	    }
    } else {
        content := form.Value["content"][0]
        artikelContent.Content = content
    } 
    
    validate := validator.New()
	
    if err := c.BodyParser(&artikelContent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(&artikelContent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": err.Error(),
        })
	}
	    
	if result := initialize.DB.Create(&artikelContent); result.Error != nil {		
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to create data artikel content",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{        
        "status": "success",        
    })
}