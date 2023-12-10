package main

import (
	"fmt"

	"github.com/dataped/mapdf"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const uploadDir = "./uploads"

func main() {
	app := fiber.New()
	// Create a directory to store uploaded files
	mapdf.CreateFolder(uploadDir)

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
		err = mapdf.SaveUploadedFile(file, uploadDir, fname)
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
