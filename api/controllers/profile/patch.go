package profile

import (
	"fmt"
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UpdateProfile struct {
	Name        string    `gorm:"size:100;not null" json:"name" validate:"required,max=100"`
	Email       string    `gorm:"size:100;unique;not null" json:"email" validate:"required,email"`		
}

func PatchProfileAdmin(c *fiber.Ctx) error {
    // Parse multipart form
    form, err := c.MultipartForm()
    if err != nil {
        return err
    }

    // Get the refresh token from cookies
    refreshToken := c.Cookies("refreshToken")

    // Parse and validate the token
    claims, err := utils.ParseToken(refreshToken, "refresh")
    if err != nil || claims["type"] != "refresh" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "status":  "error",
            "message": "You are not authenticated.",
        })
    }

    // Type-assert userID to string
    userID, ok := claims["id"].(string)
    if !ok {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status":  "error",
            "message": "Invalid user ID.",
        })
    }

    // Extract profile picture file from form
    fotoProfile := form.File["fotoProfile"]

    var admin models.Admin

    // Create new admin profile data from form values
    newAdmin := UpdateProfile{
        Name:  form.Value["name"][0],
        Email: form.Value["email"][0],
    }

    // Retrieve the current admin record
    if err := initialize.DB.Where("id = ?", userID).First(&admin).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status":  "error",
            "message": "Data admin cannot be found.",
        })
    }

    // Validate the new admin data
    validate := validator.New()
    if err := c.BodyParser(&newAdmin); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "error",
            "message": "Invalid request body",
        })
    }

    if err := validate.Struct(&newAdmin); err != nil {
        var errorMassage []string
        validationErrors := err.(validator.ValidationErrors)
        for _, fieldError := range validationErrors {
            errorMassage = append(errorMassage, utils.ErrorMassage(fieldError.Field(), fieldError.Tag(), fieldError.Param()))
        }
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "error",
            "message": errorMassage,
        })
    }

    // Update the admin profile data
    admin = models.Admin{
        Name:  newAdmin.Name,
        Email: newAdmin.Email,
    }

    // Handle profile picture upload and deletion
    if len(fotoProfile) > 0 {
        if admin.FotoProfile != nil {
            if err := utils.DeleteFile(admin.FotoProfile, initialize.ENV_DIR_ADMIN_FILES, userID); err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                    "status":  "error",
                    "message": fmt.Sprintf("Failed to delete old profile picture: %v", err),
                })
            }
        }
        uploadedFileNames, err := utils.UploadFile(fotoProfile, initialize.ENV_DIR_ADMIN_FILES, userID)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status":  "error",
                "message": fmt.Sprintf("Failed to upload new profile picture: %v", err),
            })
        }
        admin.FotoProfile = &uploadedFileNames
    }

    // Update the admin record in the database
    if result := initialize.DB.Where("id = ?", userID).Updates(&admin); result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status":  "error",
            "message": "Failed to update data admin",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status":  "success",
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

	if result := initialize.DB.Where("id = ?", userID).Updates(&admin); result.Error != nil {
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