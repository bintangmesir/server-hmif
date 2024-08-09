package utils

import "github.com/gofiber/fiber/v2"

func HttpsCheck (c *fiber.Ctx) bool {
    protocol := c.Protocol()

    if protocol == "https" {
        return true
    } else {
        return false
    }
}