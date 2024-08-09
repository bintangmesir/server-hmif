package comment

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PatchComment (c *fiber.Ctx) error {

    id := c.Params("id")

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}

    image := form.File["image"]

    var comment models.Comment
    
    newComment := models.Comment {
        Text: form.Value["text"][0],
        Email: form.Value["email"][0],        		           
    }    

    if err := initialize.DB.Where("id = ?", id).First(&comment).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data artikel cannot be found.",
		})
	}

    validate := validator.New()
	
    if err := c.BodyParser(&newComment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(&newComment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": err.Error(),
        })
	}
    
    if len(image) > 0 {		        
        if comment.Image != nil {
            if  err := utils.DeleteFile(comment.Image, initialize.ENV_DIR_COMMENT_FILES, id); err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                    "status": "error",
                    "message": "Internal server error.",
                })
		    }

        }
		uploadedFileNames, err := utils.UploadFile(image, initialize.ENV_DIR_COMMENT_FILES, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
				"message": "Internal server error.",
			})
		}
		comment.Image = &uploadedFileNames
	}

    comment = models.Comment{
        Text: newComment.Text,
        Email: newComment.Email,
    }
	
	if result := initialize.DB.Where("id = ?", id).Updates(&comment); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to update data comment",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Data artikel updated.",
    })
}