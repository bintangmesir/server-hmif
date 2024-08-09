package routes

import (
	"server/api/controllers/auth"
	"server/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute (app *fiber.App){

    api := app.Group("/api")
    v1 := api.Group("/v1")

    v1.Get("/csrf-token",auth.CSRFToken)
    v1.Get("/refresh-token", auth.RefreshTokenHandler)
    v1.Post("/login", auth.Login)
    v1.Post("/logout", middlewares.RoleMiddleware([]string{}), auth.Logout)    
}