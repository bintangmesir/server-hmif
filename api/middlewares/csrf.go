package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
)

func CSRFConfigs() csrf.Config{
    csrfConfig := csrf.Config{
        KeyLookup:      "header:X-Csrf-Token",
        CookieName:     "csrf_token",
	    CookieSameSite: "Lax",
        Expiration:     1 * time.Hour,        
        CookieDomain:   "localhost",
        KeyGenerator:   utils.UUIDv4,
    }

    return csrfConfig
}