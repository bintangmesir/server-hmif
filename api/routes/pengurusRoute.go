package routes

import (
	"server/api/controllers/pengurus"
	"server/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func PengurusRoute(app *fiber.App){    

    api := app.Group("/api")
    v1 := api.Group("/v1")
    
    v1.Get("/pengurus",middlewares.RoleMiddleware([]string{"kadep_kominfo", "staff_kominfo"}), pengurus.GetPengurus)
    v1.Get("/pengurus/:id", middlewares.RoleMiddleware([]string{"kadep_kominfo", "staff_kominfo"}), pengurus.GetByIdPengurus)
    v1.Post("/pengurus", middlewares.RoleMiddleware([]string{"staff_kominfo"}), pengurus.PostPengurus)
    v1.Patch("/pengurus/:id", middlewares.RoleMiddleware([]string{"staff_kominfo"}), pengurus.PatchPengurus)
    v1.Delete("/pengurus/:id", middlewares.RoleMiddleware([]string{"staff_kominfo"}), pengurus.DeletePengurus)
}