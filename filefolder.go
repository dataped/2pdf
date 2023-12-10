package mapdf

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func CreateFolder(pathDir string) {
	if err := os.MkdirAll(pathDir, 0755); err != nil {
		fmt.Println("Error creating upload directory:", err)
		return
	}
}

func SaveUploadedFile(file *multipart.FileHeader, uploadDir, filename string) error {
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
