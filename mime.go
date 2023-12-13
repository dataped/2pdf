package mapdf

import (
	"mime/multipart"
	"net/http"
)

func GetMimeType(file *multipart.FileHeader) (mime string, err error) {
	// Create a buffer to store the header of the file in
	fileHeader := make([]byte, 512)
	src, err := file.Open()
	if err != nil {
		return
	}
	// Copy the headers into the FileHeader buffer
	if _, err = src.Read(fileHeader); err != nil {
		return
	}
	// set position back to start.
	if _, err = src.Seek(0, 0); err != nil {
		return
	}
	mime = http.DetectContentType(fileHeader)
	return

}
