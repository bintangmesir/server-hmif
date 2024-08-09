package routes

import "github.com/gofiber/fiber/v2"

func IndexRoute (app *fiber.App){    
    AuthRoute(app)  
    ProfileRoute(app)
    CommentRoute(app)

    // Kominfo Route
    AdminRoute(app)  
    ArtikelRoute(app)
    ArtikelContentRoute(app)
    YoutubeRoute(app)
    PengurusRoute(app)

    // PRHP Route
    AlumniRoute(app)
    BarangRoute(app)
    BukuRoute(app)
}