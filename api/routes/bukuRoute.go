package routes

import (
	"server/api/controllers/buku"
	"server/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func BukuRoute(app *fiber.App){    

    api := app.Group("/api")
    v1 := api.Group("/v1")
    
    v1.Get("/buku", buku.GetBuku)

    v1.Get("/buku/:id", buku.GetByIdBuku)

    v1.Post("/buku", middlewares.RoleMiddleware([]string{"staff_prhp"}), buku.PostBuku)

    v1.Patch("/buku/:id", middlewares.RoleMiddleware([]string{"staff_prhp"}), buku.PatchBuku)
    
    v1.Delete("/buku/:id", middlewares.RoleMiddleware([]string{"staff_prhp"}), buku.DeleteBuku)
}