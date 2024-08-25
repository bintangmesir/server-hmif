package auth

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RefreshTokenHandler(c *fiber.Ctx) error {
    var admin models.Admin
    refreshToken := c.Cookies("refreshToken")

    if refreshToken == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "status": "error",
            "message": "You are not authenticated.",
        })
    }

    claims, err := utils.ParseToken(refreshToken, "refresh")
    if err != nil || claims["type"] != "refresh" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "status": "error",
            "message": "You are not authenticated.",
        })
    }

    userID := claims["id"]
    if err := initialize.DB.Where("id = ?", userID).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data admin cannot be found.",
		})
	}    

    newAccessToken, err := utils.GenerateToken(admin.ID, admin.Role, "access")
    if err != nil {        
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to generate new access token.",
        })
    }

    protocol, _ := strconv.ParseBool(initialize.ENV_POTOCOL_HTTPS)

    c.Cookie(&fiber.Cookie{
        Name:     "accessToken",
        Value:    newAccessToken,
        Expires:  time.Now().Add(time.Hour * 1), // Token expires in 1 hour                       
        HTTPOnly: false,        
        Secure:   protocol,       
    })
    
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "status": "success",                
    })
}