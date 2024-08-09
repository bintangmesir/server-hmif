package main

import (
	"server/api/middlewares"
	"server/api/routes"
	"server/initialize"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main(){
    initialize.EnvVariables()
    initialize.DBConnection()

    app := fiber.New()
    app.Use(cors.New(middlewares.CORSConfigs()))
    app.Use(csrf.New(middlewares.CSRFConfigs()))
    app.Use(limiter.New(middlewares.LimiterConfigs()))
    app.Use(compress.New(compress.Config{
        Level: compress.LevelBestSpeed,
    }))
    app.Use(recover.New())
    app.Static("/static", "./public")

    routes.IndexRoute(app)

    app.Listen(":" + initialize.ENV_PORT)
}