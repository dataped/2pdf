package main

import (
	"fmt"
)

func main() {
	// Example usage
	targetURL := "https://komarbe.ulbi.ac.id/files/upload"
	tokenKey := "Login"
	tokenValue := "v4.public.eyJkYXRhIjp7ImlkX3BlbmRhZnRhciI6NzYsImlkX2Jpb2RhdGEiOiJjbHJ2dGM5dTg0M21kYms4ZWJiZyIsInBob25lX251bWJlciI6IjYyODU3MjI2OTc5MTgiLCJlbWFpbCI6ImNocmlzeXVkYUB1bGJpLmFjLmlkIiwiZnVsbF9uYW1lIjoiY2hyaXN0aWFuIHl1ZGEgcHJhdGFtYSJ9LCJleHAiOiIyMDIzLTEyLTE0VDA4OjQwOjI2KzA3OjAwIiwiaWF0IjoiMjAyMy0xMi0xM1QxNDo0MDoyNiswNzowMCIsImlkIjoiNjI4NTcyMjY5NzkxOCIsIm5iZiI6IjIwMjMtMTItMTNUMTQ6NDA6MjYrMDc6MDAifZMuLrRMEVB4gE_GOe8hLo6NrE_YGPn3NrIceJK_L0oSHtnlDzFlLKZXMf94I23k1ZGMlzl0VMhnRgKJrEt8eAg"
	filePath := "./image1.jpeg"
	formdataName := "file"

	responseBody, err := mapdf.postFileWithHeader(targetURL, tokenKey, tokenValue, filePath, formdataName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", string(responseBody))
}

// const uploadDir = "./uploads"

// func main() {
// 	app := fiber.New()
// 	// Create a directory to store uploaded files
// 	mapdf.CreateFolder(uploadDir)

// 	// Endpoint to upload an image
// 	app.Post("/upload", func(ctx *fiber.Ctx) error {
// 		// Parse the form file
// 		file, err := ctx.FormFile("image")
// 		if err != nil {
// 			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
// 		}

// 		// Save the uploaded file to the server
// 		id := uuid.New()
// 		fname := id.String() + ".jpg"
// 		err = mapdf.SaveUploadedFile(file, uploadDir, fname)
// 		if err != nil {
// 			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 		}

// 		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"filename": fname})
// 	})

// 	// Start the Fiber app on port 3000
// 	err := app.Listen(":3000")
// 	if err != nil {
// 		fmt.Println("Error starting the server:", err)
// 	}
// }
