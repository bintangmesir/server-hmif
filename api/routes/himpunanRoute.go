package routes

import (
	"server/api/controllers/himpunan"
	"server/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func HimpunanRoute(app *fiber.App){    

    api := app.Group("/api")
    v1 := api.Group("/v1")
    
    v1.Get("/himpunan", himpunan.GetHimpunan)

    v1.Patch("/himpunan/:id", middlewares.RoleMiddleware([]string{"kadep_kominfo", "staff_kominfo"}), himpunan.PatchHimpunan)
}