package auth

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

    if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": err.Error(),
        })
	}

    if err := initialize.DB.Where("email = ?", req.Email).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data admin cannot be found.",
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

    c.Cookie(&fiber.Cookie{
        Name:     "accessToken",
        Value:    accessToken,
        Expires:  time.Now().Add(time.Hour * 1), // Token expires in 1 hour                       
        HTTPOnly: false,        
        Secure:   utils.HttpsCheck(c),       
    })

    c.Cookie(&fiber.Cookie{
        Name:     "refreshToken",
        Value:    refreshToken,
        Expires:  time.Now().Add(time.Hour * 24 * 7), // Refresh token expires in 7 days
        HTTPOnly: true,        
        Secure:   utils.HttpsCheck(c),        
    })

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status":"error",
        "message": "Logged in successfully",
    })   
}
