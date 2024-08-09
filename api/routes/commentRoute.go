package routes

import (
	"server/api/controllers/comment"
	"server/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func CommentRoute(app *fiber.App){    

    api := app.Group("/api")
    v1 := api.Group("/v1")
    
    v1.Post("/comment", middlewares.RoleMiddleware([]string{}), comment.PostComment)
    v1.Patch("/comment/:id", middlewares.RoleMiddleware([]string{}), comment.PatchComment)
    v1.Delete("/comment/:id", middlewares.RoleMiddleware([]string{}), comment.DeleteComment)
}