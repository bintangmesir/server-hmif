package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func LimiterConfigs () limiter.Config {
    limiterConfig := limiter.Config{
        Next: func(c *fiber.Ctx) bool {
            return c.IP() == "127.0.0.1"
        },
        Max:          120,
        Expiration:     60 * time.Second,
        KeyGenerator:          func(c *fiber.Ctx) string {
            return c.Get("x-forwarded-for")
        },
        LimitReached: func(c *fiber.Ctx) error {
            return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
                "status" : "error",
                "message" : "Too many request.",
            })
        },        
    }
    return limiterConfig
}