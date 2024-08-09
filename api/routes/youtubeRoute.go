package routes

import (
	"server/api/controllers/youtube"
	"server/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func YoutubeRoute(app *fiber.App){    

    api := app.Group("/api")
    v1 := api.Group("/v1")
    
    v1.Get("/youtube",middlewares.RoleMiddleware([]string{"kadep_kominfo", "staff_kominfo"}), youtube.GetYoutube)
    v1.Get("/youtube/:id", middlewares.RoleMiddleware([]string{"kadep_kominfo", "staff_kominfo"}), youtube.GetByIdYoutube)
    v1.Post("/youtube", middlewares.RoleMiddleware([]string{"staff_kominfo"}), youtube.PostYoutube)
    v1.Patch("/youtube/:id", middlewares.RoleMiddleware([]string{"staff_kominfo"}), youtube.PatchYoutube)
    v1.Delete("/youtube/:id", middlewares.RoleMiddleware([]string{"staff_kominfo"}), youtube.DeleteYoutube)
}