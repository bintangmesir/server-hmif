package buku

import (
	"server/initialize"
	"server/internal/models"
	"server/pkg/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PatchBuku (c *fiber.Ctx) error {

    id := c.Params("id")

    form, err := c.MultipartForm()
	if err != nil {
		return err
	}

    cover := form.File["cover"]

    jumlah, err := strconv.Atoi(form.Value["jumlah"][0]); if err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to convert variable jumlah"})
    }

    var buku models.Buku

    newBuku := models.Buku {
        Judul: form.Value["judul"][0],
        Kode: form.Value["kode"][0],
        Penulis: form.Value["penulis"][0],        
		TahunTerbit: form.Value["tahunTerbit"][0],
        Penerbit: form.Value["penerbit"][0],
        Abstrak: form.Value["abstrak"][0],
        Jumlah: int64(jumlah),
    }    

    if err := initialize.DB.Where("id = ?", id).First(&buku).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": "error",
			"message": "Data buku cannot be found.",
		})
	}

    validate := validator.New()
	
    if err := c.BodyParser(&newBuku); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": "Invalid request body",
        })
	}

	if err := validate.Struct(newBuku); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": "error",
            "message": err.Error(),
        })
	}
    
    if len(cover) > 0 {		        
        if buku.Cover != nil {
            if  err := utils.DeleteFile(buku.Cover, initialize.ENV_DIR_BUKU_FILES, id); err != nil {
                return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                    "status": "error",
                    "message": "Internal server error.",
                })
		    }

        }
		uploadedFileNames, err := utils.UploadFile(cover, initialize.ENV_DIR_BUKU_FILES, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
				"message": "Internal server error.",
			})
		}
		buku.Cover = &uploadedFileNames
	}

    buku = models.Buku{
        Judul: newBuku.Judul,
        Kode: newBuku.Kode,
        Penulis: newBuku.Penulis,        
		TahunTerbit: newBuku.TahunTerbit,
        Penerbit: newBuku.Penerbit,
        Abstrak: newBuku.Abstrak,
        Jumlah: newBuku.Jumlah,        
    }
	
	if result := initialize.DB.Where("id = ?", id).Updates(&buku); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": "error",
            "message": "Failed to update data buku",
        })
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "status": "success",
        "message": "Data buku updated.",
    })
}