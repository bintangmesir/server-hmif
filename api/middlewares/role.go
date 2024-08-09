package middlewares

import (
	"server/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func RoleMiddleware(roles []string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Get the access token from the header
        tokenStr := c.Get("Authorization")
        if tokenStr == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "status":"error",
                "message": "You are not authenticated.",
            })
        }

        // Remove "Bearer " prefix if present
        if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
            tokenStr = tokenStr[7:]
        }

        claims, err := utils.ParseToken(tokenStr, "access")
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "status":"error",
                "message": "You are not authenticated.",
            })
        }

        // Role validation
        if len(roles) > 0 {
            role, ok := claims["role"].(string)
            if !ok || !contains(roles, role) {
                return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
                    "status":"error",
                    "message": "Insufficient role",
                })
            }
        }

        return c.Next()
    }
}

func contains(slice []string, item string) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}