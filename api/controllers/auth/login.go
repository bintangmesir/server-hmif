package auth

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}

func Login(c *fiber.Ctx) error {    

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}    
    
    req := LoginRequest{
        Email: form.Value["email"][0],
        Password: form.Value["password"][0],
    }

    var admin models.Admin
    validate := validator.New()

    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request",
        })
    }

    if err := validate.Struct(&req); err != nil {
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

    if err := initialize.DB.Where("email = ?", req.Email).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data admin cannot be found.",
		})
	}
    
    if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Password is incorrect.",
		})
	}

    accessToken, err := utils.GenerateToken(admin.ID, admin.Role, "access")
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to generate token",
        })
    }

    refreshToken, err := utils.GenerateToken(admin.ID, admin.Role, "refresh")
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to generate token",
        })
    }

    protocol, _ :=strconv.ParseBool(initialize.ENV_POTOCOL_HTTPS)

    c.Cookie(&fiber.Cookie{
        Name:     "accessToken",
        Value:    accessToken,
        Expires:  time.Now().Add(time.Hour * 1), // Token expires in 1 hour                       
        HTTPOnly: false,        
        Secure:   protocol,       
    })

    c.Cookie(&fiber.Cookie{
        Name:     "refreshToken",
        Value:    refreshToken,
        Expires:  time.Now().Add(time.Hour * 24 * 7), // Refresh token expires in 7 days
        HTTPOnly: true,        
        Secure:   protocol,        
    })

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status":"success",
        "message": "Logged in successfully",
    })   
}
