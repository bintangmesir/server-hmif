package alumni

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

func GetAlumni(c *fiber.Ctx) error {
    var params PaginationParams
    if err := c.QueryParser(&params); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "": "Invalid query parameters",
        })
    }
    
    if params.Limit == 0 {
        params.Limit = 10
    }

    var totalCount int64
    var alumni []models.Alumni
    
    if err := initialize.DB.Model(&models.Alumni{}).Count(&totalCount).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to count alumni",
        })
    }
    
    query := initialize.DB.Model(&models.Alumni{})

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
    
    query = query.Order("id").Limit(params.Limit).Find(&alumni)
    if query.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to retrieve alumni",
        })
    }
    
    var nextCursor string
    if len(alumni) > 0 {
        nextCursor = alumni[len(alumni)-1].ID.String()
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "count":      totalCount,
        "next":       nextCursor,
        "limit":      params.Limit,
        "data":       alumni,
    })
}