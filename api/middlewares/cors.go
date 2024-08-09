package middlewares

import (
	"server/initialize"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORSConfigs () cors.Config{
     corsConfig := cors.Config{
        AllowOrigins: initialize.ENV_CLIENT_URI,
        AllowMethods: "GET, POST, PUT, DELETE, OPTIONS, PATCH",
        AllowHeaders: "Origin, Content-Type, Accept, Authorization, X-Requested-With, X-Csrf-Token",
        AllowCredentials: true,
    }

    return corsConfig
}