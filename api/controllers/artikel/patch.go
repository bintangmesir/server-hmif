package artikel

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PatchArtikel (c *fiber.Ctx) error {

    id := c.Params("id")

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}

    thumbnail := form.File["thumbnail"]

    commentEnabled, err := strconv.ParseBool(form.Value["commentEnabled"][0]); if err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to convert Comment Enabled variable ",
        })
    }

    var artikel models.Artikel

    newArtikel := models.Artikel {        
        Title: form.Value["title"][0],
        SubTitle: form.Value["subTitle"][0],
        CommentEnabled: commentEnabled,        		
    }    
    
    if err := initialize.DB.Where("id = ?", id).First(&artikel).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data artikel cannot be found.",
		})
	}

    validate := validator.New()
	
    if err := c.BodyParser(&newArtikel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(&newArtikel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": err.Error(),
        })
	}
    
    if len(thumbnail) > 0 {		        
        if artikel.Thumbnail != nil {
            if  err := utils.DeleteFile(artikel.Thumbnail, initialize.ENV_DIR_ARTIKEL_FILES, id); err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                    "status": "error",
                    "message": "Internal server error.",
                })
		    }

        }
		uploadedFileNames, err := utils.UploadFile(thumbnail, initialize.ENV_DIR_ARTIKEL_FILES, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
				"message": "Internal server error.",
			})
		}
		artikel.Thumbnail = &uploadedFileNames
	}

    artikel = models.Artikel{
        Title: newArtikel.Title,
        SubTitle: newArtikel.SubTitle,
        CommentEnabled: newArtikel.CommentEnabled,        		        
    }
	
	if result := initialize.DB.Where("id = ?", id).Updates(&artikel); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to update data artikel",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Data artikel updated.",
    })
}

func PatchArtikelView (c *fiber.Ctx) error {
    id := c.Params("id")

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}

    var artikel models.Artikel

    view, err := strconv.Atoi(form.Value["view"][0]); if err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to convert variable jumlah"})
    }    
    
    if err := initialize.DB.Where("id = ?", id).First(&artikel).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data artikel cannot be found.",
		})
	}

    artikel.View = int64(view)

    if result := initialize.DB.Where("id = ?", id).Updates(&artikel); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to update data artikel",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Data artikel updated.",
    })
}