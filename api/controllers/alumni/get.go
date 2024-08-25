package alumni

import (
	"server/initialize"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
)

type PaginationParams struct {
    Offset int `query:"offset"`
    Limit  int `query:"limit"`
}

func GetAlumni(c *fiber.Ctx) error {
    var params PaginationParams
    if err := c.QueryParser(&params); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid query parameters",
        })
    }

    var totalCount int64
    var alumni []models.Alumni
    
    // Count the total number of records
    if err := initialize.DB.Model(&models.Alumni{}).Count(&totalCount).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to count alumni",
        })
    }
    
    if params.Limit == 0 {
        params.Limit = int(totalCount)
    }

    // Offset default to 0 if not provided or negative
    if params.Offset < 0 {
        params.Offset = 0
    }

    query := initialize.DB.Model(&models.Alumni{})

    // Apply offset and limit
    query = query.Order("created_at DESC").Offset(params.Offset).Limit(params.Limit).Find(&alumni)
    
    if query.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to retrieve alumni",
        })
    }

    // Calculate the next offset
    offset := params.Offset + params.Limit
    if offset > int(totalCount) {
        offset = int(totalCount)
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "count":      totalCount,
        "offset": offset,
        "limit":      params.Limit,
        "data":       alumni,
    })
}
