package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadFile(filenames []*multipart.FileHeader, filePath, id string) (string, error) {	
	cwd, err := os.Getwd()
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Failed to get current working directory")
	}

	// Define the directory path with the given ID
	dirPath := fmt.Sprintf("%s%s%s/", cwd, filePath, id)

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Failed to create directory")
	}

	var uploadedFileNames []string
	for _, fileHeader := range filenames {		
		if fileHeader.Size > 1e6 { // 1 MB limit
			return "", fiber.NewError(fiber.StatusBadRequest, "File size exceeds 1 MB")
		}

		// Open the file
		src, err := fileHeader.Open()
		if err != nil {
			return "", fiber.NewError(fiber.StatusInternalServerError, "Failed to open file")
		}
		defer src.Close()

		buffer := make([]byte, 512)
		if _, err := src.Read(buffer); err != nil {
			return "", fiber.NewError(fiber.StatusInternalServerError, "Failed to read file")
		}
		src.Seek(0, io.SeekStart)

		fileType := http.DetectContentType(buffer)
		if !(fileType == "image/jpeg" || fileType == "image/jpg" || fileType == "image/png") {
			return "", fiber.NewError(fiber.StatusBadRequest, "Invalid file type. Only JPEG, JPG, and PNG are allowed.")
		}

		newFilename := time.Now().Format("20060102150405") + "_" + fileHeader.Filename
		dstPath := fmt.Sprintf("%s%s", dirPath, newFilename)

		dst, err := os.Create(dstPath)
		if err != nil {
			return "", fiber.NewError(fiber.StatusInternalServerError, "Failed to create destination file")
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return "", fiber.NewError(fiber.StatusInternalServerError, "Failed to copy file")
		}

		uploadedFileNames = append(uploadedFileNames, newFilename)
	}

	return strings.Join(uploadedFileNames, ";"), nil
}

func DeleteFile(filenamesString *string, filePath, id string) error {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Define the directory path with the given ID
	dirPath := fmt.Sprintf("%s%s%s/", cwd, filePath, id)

	// Get the list of files to delete
	filenames := strings.Split(*filenamesString, ";")

	for _, filename := range filenames {
		filePath := fmt.Sprintf("%s%s", dirPath, filename)

		if err := os.Remove(filePath); err != nil {
			fmt.Println(err)
			return err
		}
	}

	// Remove the directory if empty
	if err := os.Remove(dirPath); err != nil {		
		return err
	}

	return nil
}
