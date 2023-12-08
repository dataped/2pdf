package main

import (
	"fmt"
	_ "image/jpeg" // Importing the image/jpeg package for JPEG decoding

	"github.com/jung-kurt/gofpdf"
)

func main() {
	outputPDF := "merged_images.pdf"
	imageFiles := []string{"image1.jpeg", "image2.jpg", "image3.jpeg"} // Replace with your image file paths

	if err := mergeImagesToPDF(outputPDF, imageFiles); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("PDF created successfully:", outputPDF)
	}
}

func mergeImagesToPDF(outputPDF string, imageFiles []string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	for _, imageFile := range imageFiles {
		pdf.ImageOptions(imageFile, 10, 10, 0, 0, false, gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true}, 0, "")
		pdf.AddPage()
	}

	return pdf.OutputFileAndClose(outputPDF)
}
