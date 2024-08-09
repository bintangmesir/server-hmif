package admin

import (
	"server/initialize"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PaginationParams struct {
    Cursor string `query:"cursor"`
    Limit  int    `query:"limit"`
}

func GetAdmins(c *fiber.Ctx) error {
    var params PaginationParams
    if err := c.QueryParser(&params); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid query parameters",
        })
    }
    
    if params.Limit == 0 {
        params.Limit = 10
    }

    var totalCount int64
    var admins []models.Admin
    
    if err := initialize.DB.Model(&models.Admin{}).Count(&totalCount).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to count users",
        })
    }
    
    query := initialize.DB.Model(&models.Admin{})

    if params.Cursor != "" {        
        cursorUUID, err := uuid.Parse(params.Cursor)
        if err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "status": "error",
                "message": "Invalid cursor UUID",
            })
        }
        query = query.Where("id > ?", cursorUUID)
    }
    
    query = query.Order("id").Limit(params.Limit).Find(&admins)
    if query.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to retrieve users",
        })
    }
    
    var nextCursor string
    if len(admins) > 0 {
        nextCursor = admins[len(admins)-1].ID.String()
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "count":    totalCount,
        "next":     nextCursor,
        "limit":    params.Limit,
        "data":     admins,
    })
}