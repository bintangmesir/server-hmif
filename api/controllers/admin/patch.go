package admin

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type PatchAdminBody struct {
	Name        string    `gorm:"size:100;not null" json:"name" validate:"required,max=100"`
	Email       string    `gorm:"size:100;unique;not null" json:"email" validate:"required,email"`
	FotoProfile *string   `gorm:"size:255" json:"fotoProfile"`
}

func PatchAdmin (c *fiber.Ctx) error {

    id := c.Params("id")

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	var admin models.Admin
    
	fotoProfile := form.File["fotoProfile"]
        
    newAdmin := PatchAdminBody {
        Name: form.Value["name"][0],
        Email: form.Value["email"][0],                
    }    

    if err := initialize.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data admin cannot be found.",
		})
	}

	 if err := initialize.DB.Where("email = ? AND id != ?", newAdmin.Email, id).First(&admin).Error; err == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Email admin already used.",
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
	
	admin = models.Admin{
		Name: newAdmin.Name,
		Email: newAdmin.Email,
		FotoProfile: newAdmin.FotoProfile,
	}

	if len(fotoProfile) > 0 {		        
        if admin.FotoProfile != nil {
            if  err := utils.DeleteFile(admin.FotoProfile, initialize.ENV_DIR_ADMIN_FILES, id); err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                    "status": "error",
                    "message": "Internal server error.",
                })
		    }

        }
		uploadedFileNames, err := utils.UploadFile(fotoProfile, initialize.ENV_DIR_ADMIN_FILES, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
				"message": "Internal server error.",
			})
		}
		admin.FotoProfile = &uploadedFileNames
	}

	if result := initialize.DB.Where("id = ?", id).Updates(&admin); result.Error != nil {
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

func UpdatePassword(c *fiber.Ctx) error {
	id := c.Params("id")

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	oldPassword := form.Value["oldPassword"][0]
	newPassword := form.Value["newPassword"][0]
	
	var admin models.Admin
	if err := initialize.DB.Where("id = ?", id).First(&admin).Error; err != nil {
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
	if result := initialize.DB.Where("id = ?", id).Updates(&admin); result.Error != nil {
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

func ResetPassword(c *fiber.Ctx) error {
	id := c.Params("id")

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	
	password := form.Value["password"][0]
	
	var admin models.Admin
	if err := initialize.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Admin not found.",
		})
	}
	
	if len(password) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "New password must be at least 6 characters long.",
		})
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to hash new password.",
		})
	}
	
	admin.Password = string(hashedPassword)
	if result := initialize.DB.Where("id = ?", id).Updates(&admin); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to reset password.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Password reseted successfully.",
	})
}