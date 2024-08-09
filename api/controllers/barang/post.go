package barang

import (
	"server/initialize"
	"server/internal/models"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostBarang (c *fiber.Ctx) error {

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}    

    jumlah, err := strconv.Atoi(form.Value["jumlah"][0]); if err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to convert variable jumlah"})
    }

    baik, err := strconv.Atoi(form.Value["baik"][0]); if err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to convert variable baik"})
    }

    rusakRingan, err := strconv.Atoi(form.Value["rusakRingan"][0]); if err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to convert variable Rusak Ringan"})
    }

    rusakBerat, err := strconv.Atoi(form.Value["rusakBerat"][0]); if err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to convert variable Rusak Berat"})
    }
    
    barang := models.Barang {
        Nama: form.Value["nama"][0],                
        Jumlah: int64(jumlah),
        Baik: int64(baik),                
        RusakRingan: int64(rusakRingan),                
        RusakBerat: int64(rusakBerat),
        Keterangan: form.Value["keterangan"][0],                
    }    

    if uuidStr, err := uuid.NewUUID(); err == nil {
		barang.ID = uuidStr
	} else {
		return err
	}  

    validate := validator.New()
	
    if err := c.BodyParser(&barang); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(barang); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": err.Error(),
        })
	}

	if result := initialize.DB.Create(&barang); result.Error != nil {		
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to create data barang",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Data barang created.",
    })
}