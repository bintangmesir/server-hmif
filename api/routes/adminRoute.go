package routes

import (
	"server/api/controllers/admin"
	"server/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AdminRoute(app *fiber.App){    

    api := app.Group("/api")
    v1 := api.Group("/v1")
    
    v1.Get("/admin",middlewares.RoleMiddleware([]string{"super_admin","kadep_kominfo", "kadep_prhp"}), admin.GetAdmins)

    v1.Get("/admin/:id", middlewares.RoleMiddleware([]string{}), admin.GetByIdAdmin)

    v1.Post("/admin", middlewares.RoleMiddleware([]string{"super_admin","kadep_kominfo", "kadep_prhp"}), admin.PostAdmin)

    v1.Patch("/admin/:id", middlewares.RoleMiddleware([]string{"super_admin","kadep_kominfo", "kadep_prhp"}), admin.PatchAdmin)

    v1.Patch("/admin/update-password/:id", middlewares.RoleMiddleware([]string{"super_admin","kadep_kominfo", "kadep_prhp"}), admin.UpdatePassword)

    v1.Patch("/admin/reset-password/:id", middlewares.RoleMiddleware([]string{"super_admin"}), admin.ResetPassword)
    
    v1.Delete("/admin/:id", middlewares.RoleMiddleware([]string{"super_admin","kadep_kominfo", "kadep_prhp"}), admin.DeleteAdmin)
}