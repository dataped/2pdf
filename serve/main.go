package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const uploadDir = "./uploads"

func main() {
	app := fiber.New()

	// Create a directory to store uploaded files
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		fmt.Println("Error creating upload directory:", err)
		return
	}

	// Endpoint to upload an image
	app.Post("/upload", func(ctx *fiber.Ctx) error {
		// Parse the form file
		file, err := ctx.FormFile("image")
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// Save the uploaded file to the server
		id := uuid.New()
		fname := id.String() + ".jpg"
		err = saveUploadedFile(file, fname)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"filename": fname})
	})

	// Start the Fiber app on port 3000
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

func saveUploadedFile(file *multipart.FileHeader, filename string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Create a new file in the upload directory
	dst, err := os.Create(fmt.Sprintf("%s/%s", uploadDir, filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the contents of the uploaded file to the new file
	_, err = io.Copy(dst, src)
	return err
}
