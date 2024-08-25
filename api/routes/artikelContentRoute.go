package routes

import (
	artikelcontent "server/api/controllers/artikel-content"
	"server/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ArtikelContentRoute(app *fiber.App){    

    api := app.Group("/api")
    v1 := api.Group("/v1")
    
    v1.Post("/artikel-content", middlewares.RoleMiddleware([]string{"staff_kominfo"}), artikelcontent.PostArtikelContent)
}