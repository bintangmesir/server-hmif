package routes

import (
	"server/api/controllers/barang"
	"server/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func BarangRoute(app *fiber.App){    

    api := app.Group("/api")
    v1 := api.Group("/v1")
    
    v1.Get("/barang", middlewares.RoleMiddleware([]string{"kadep_prhp", "staff_prhp"}),barang.GetBarang)
    v1.Get("/barang/:id", middlewares.RoleMiddleware([]string{"kadep_prhp", "staff_prhp"}), barang.GetByIdBarang)
    v1.Post("/barang",middlewares.RoleMiddleware([]string{"staff_prhp"}), barang.PostBarang)
    v1.Patch("/barang/:id",middlewares.RoleMiddleware([]string{"staff_prhp"}), barang.PatchBarang)
    v1.Delete("/barang/:id",middlewares.RoleMiddleware([]string{"staff_prhp"}), barang.DeleteBarang)
}