package routes

import (
	"server/api/controllers/alumni"
	"server/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AlumniRoute(app *fiber.App){    

    api := app.Group("/api")
    v1 := api.Group("/v1")
    
    v1.Get("/alumni", middlewares.RoleMiddleware([]string{"kadep_prhp", "staff_prhp"}),alumni.GetAlumni)
    v1.Get("/alumni/:id", middlewares.RoleMiddleware([]string{"kadep_prhp", "staff_prhp"}), alumni.GetByIdAlumni)
    v1.Post("/alumni",middlewares.RoleMiddleware([]string{"staff_prhp"}), alumni.PostAlumni)
    v1.Patch("/alumni/:id",middlewares.RoleMiddleware([]string{"staff_prhp"}), alumni.PatchAlumni)
    v1.Delete("/alumni/:id",middlewares.RoleMiddleware([]string{"staff_prhp"}), alumni.DeleteAlumni)
}