package routes

import (
	"server/api/controllers/profile"
	"server/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ProfileRoute(app *fiber.App){    

    api := app.Group("/api")
    v1 := api.Group("/v1")
    
    v1.Patch("/profile", middlewares.RoleMiddleware([]string{}), profile.PatchProfileAdmin)
    v1.Patch("/profile/update-password", middlewares.RoleMiddleware([]string{}), profile.UpdatePasswordProfileAdmin)
}