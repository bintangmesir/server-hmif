package auth

import (
	"server/initialize"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {    

    protocol, _ :=strconv.ParseBool(initialize.ENV_POTOCOL_HTTPS)
   
     c.Cookie(&fiber.Cookie{
        Name:     "accessToken",
        Value:    "",
        Expires:  time.Now().Add(-time.Hour),        
        Secure:   protocol,        
    })

    c.Cookie(&fiber.Cookie{
        Name:     "refreshToken",
        Value:    "",
        Expires:  time.Now().Add(-time.Hour),
        HTTPOnly: true,
        Secure:   protocol,        
    })

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Logout success.",
    })
}