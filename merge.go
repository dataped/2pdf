package mapdf

import (
	"fmt"

	_ "image/gif"
	_ "image/jpeg" // Importing the image/jpeg package for JPEG decoding
	_ "image/png"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
)

func MergeFiles(uploadDir string, imageFiles []string) string {
	outputPDF := uploadDir + "/" + uuid.New().String() + ".pdf"
	//imageFiles := []string{"image1.jpeg", "image2.jpg", "image3.jpeg"} // Replace with your image file paths

	if err := mergeImagesToPDF(outputPDF, imageFiles); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("PDF created successfully:", outputPDF)
	}
	return outputPDF
}

func mergeImagesToPDF(outputPDF string, imageFiles []string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	for _, imageFile := range imageFiles {
		pdf.ImageOptions(imageFile, 10, 10, 0, 0, false, gofpdf.ImageOptions{ReadDpi: true}, 0, "")
		pdf.AddPage()
	}

	return pdf.OutputFileAndClose(outputPDF)
}
