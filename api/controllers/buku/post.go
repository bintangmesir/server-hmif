package buku

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostBuku (c *fiber.Ctx) error {

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}

    cover := form.File["cover"]

    jumlah, err := strconv.Atoi(form.Value["jumlah"][0]); if err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to convert variable jumlah"})
    }

    buku := models.Buku {
        Judul: form.Value["judul"][0],
        Kode: form.Value["kode"][0],
        Penulis: form.Value["penulis"][0],        
		TahunTerbit: form.Value["tahunTerbit"][0],
        Penerbit: form.Value["penerbit"][0],
        Abstrak: form.Value["abstrak"][0],
        Jumlah: int64(jumlah),
    }    

    if uuidStr, err := uuid.NewUUID(); err == nil {
		buku.ID = uuidStr
	} else {
		return err
	}    

    validate := validator.New()
	
    if err := c.BodyParser(&buku); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(buku); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": err.Error(),
        })
	}

	 if len(cover) > 0 {			
		uploadedFileNames, err := utils.UploadFile(cover, initialize.ENV_DIR_BUKU_FILES, buku.ID.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
				"message": "Internal server error.",
			})
		}
		buku.Cover = &uploadedFileNames
	}
	
	if result := initialize.DB.Create(&buku); result.Error != nil {		
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to create data buku",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Data buku created.",
    })
}