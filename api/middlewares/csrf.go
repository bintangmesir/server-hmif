package middlewares

import (
	"server/initialize"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
)

func CSRFConfigs() csrf.Config{
    protocol, _ :=strconv.ParseBool(initialize.ENV_POTOCOL_HTTPS)
    
    csrfConfig := csrf.Config{
        KeyLookup:      "header:X-Csrf-Token",
        CookieName:     "csrf_token",
        Expiration:     1 * time.Hour,        
        CookieSameSite: "None", 
        CookieSecure:   protocol,
        CookieDomain:   "localhost",
        KeyGenerator:   utils.UUIDv4,
    }

    return csrfConfig
}