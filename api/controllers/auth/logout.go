package auth

import (
	"server/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {    
   
     c.Cookie(&fiber.Cookie{
        Name:     "accessToken",
        Value:    "",
        Expires:  time.Now().Add(-time.Hour),        
        Secure:   utils.HttpsCheck(c),        
    })

    c.Cookie(&fiber.Cookie{
        Name:     "refreshToken",
        Value:    "",
        Expires:  time.Now().Add(-time.Hour),
        HTTPOnly: true,
        Secure:   utils.HttpsCheck(c),        
    })

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Logout success.",
    })
}