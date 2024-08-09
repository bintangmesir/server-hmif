package routes

import (
	"server/api/controllers/artikel"
	"server/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ArtikelRoute(app *fiber.App){    

    api := app.Group("/api")
    v1 := api.Group("/v1")
    
    v1.Get("/artikel",middlewares.RoleMiddleware([]string{"kadep_kominfo", "staff_kominfo"}), artikel.GetArtikel)
    v1.Get("/artikel/:id", middlewares.RoleMiddleware([]string{"kadep_kominfo", "staff_kominfo"}), artikel.GetByIdArtikel)
    v1.Post("/artikel", middlewares.RoleMiddleware([]string{"staff_kominfo"}), artikel.PostArtikel)    
    v1.Patch("/artikel/:id", middlewares.RoleMiddleware([]string{"staff_kominfo"}), artikel.PatchArtikel)
    v1.Patch("/artikel/:id/view", artikel.PatchArtikelView)
    v1.Delete("/artikel/:id", middlewares.RoleMiddleware([]string{"staff_kominfo"}), artikel.DeleteArtikel)
}