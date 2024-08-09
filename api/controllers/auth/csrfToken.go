package auth

import (
	"github.com/gofiber/fiber/v2"
)

func CSRFToken (c *fiber.Ctx) error{
   token := c.Cookies("csrf_token")

    if token == "" {        
        return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "status": "success",        
    })
}