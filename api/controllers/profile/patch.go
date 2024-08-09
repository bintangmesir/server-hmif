package profile

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func PatchProfileAdmin (c *fiber.Ctx) error {   
    form, err := c.MultipartForm()
	if err != nil {
		return err
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
    fotoProfile := form.File["fotoProfile"]
		
	var admin models.Admin

    newAdmin := models.Admin {
        Name: form.Value["name"][0],
        Email: form.Value["email"][0],                
    }    

    if err := initialize.DB.Where("id = ?", userID).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data admin cannot be found.",
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": err.Error(),
        })
	}    
	
    if len(fotoProfile) > 0 {		        
        if admin.FotoProfile != nil {
            if  err := utils.DeleteFile(admin.FotoProfile, initialize.ENV_DIR_ADMIN_FILES, string(admin.ID.String())); err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                    "status": "error",
                    "message": "Internal server error.",
                })
		    }

        }		
		uploadedFileNames, err := utils.UploadFile(fotoProfile, initialize.ENV_DIR_ADMIN_FILES, string(admin.ID.String()))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
				"message": "Internal server error.",
			})
		}
		admin.FotoProfile = &uploadedFileNames
	}
	
	if result := initialize.DB.Where("id = ?", userID).Updates(&admin); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to update data admin",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Data admin updated.",
    })
}

func UpdatePasswordProfileAdmin(c *fiber.Ctx) error {	

	form, err := c.MultipartForm()
	if err != nil {
		return err
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

	oldPassword := form.Value["oldPassword"][0]
	newPassword := form.Value["newPassword"][0]
	
	var admin models.Admin
	if err := initialize.DB.Where("id = ?", userID).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Admin not found.",
		})
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(oldPassword)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Old password is incorrect.",
		})
	}
	
	if len(newPassword) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "New password must be at least 6 characters long.",
		})
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to hash new password.",
		})
	}

	// Update the password
	admin.Password = string(hashedPassword)
	if result := initialize.DB.Updates(&admin); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to update password.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Password updated successfully.",
	})
}