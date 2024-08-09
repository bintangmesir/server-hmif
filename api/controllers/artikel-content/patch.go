package artikelcontent

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)


func PatchArtikelContent (c *fiber.Ctx) error {

    id := c.Params("id")

    var artikelContent models.ArtikelContent

    if err := initialize.DB.Where("id = ?", id).First(&artikelContent).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data artikel content cannot be found.",
		})
	}

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}

     index, err := strconv.Atoi(form.Value["index"][0]); if err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status":"error",
            "message": "Failed to convert variable index",
        })
    }

    newArtikelContent := models.ArtikelContent {
        Index: int64(index),
        Tipe: form.Value["tipe"][0],
        SubTipe: form.Value["subTipe"][0],        
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

    if err := c.BodyParser(&newArtikelContent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(&newArtikelContent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": err.Error(),
        })
	}

    artikelContent = models.ArtikelContent{
        Index: newArtikelContent.Index,
        Tipe: newArtikelContent.Tipe,
        SubTipe: newArtikelContent.SubTipe,
    }
	    
	if result := initialize.DB.Where("id = ?", id).Updates(&artikelContent); result.Error != nil {		
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to update data artikel content",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{        
        "status": "success",        
    })
}