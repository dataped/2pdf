package mapdf

import (
	"fmt"
	"testing"
)

func TestSendFile(t *testing.T) {
	targetURL := "https://komarbe.ulbi.ac.id/files/upload"
	tokenKey := "Login"
	tokenValue := "v4.public.eyJkYXRhIjp7ImlkX3BlbmRhZnRhciI6NzYsImlkX2Jpb2RhdGEiOiJjbHJ2dGM5dTg0M21kYms4ZWJiZyIsInBob25lX251bWJlciI6IjYyODU3MjI2OTc5MTgiLCJlbWFpbCI6ImNocmlzeXVkYUB1bGJpLmFjLmlkIiwiZnVsbF9uYW1lIjoiY2hyaXN0aWFuIHl1ZGEgcHJhdGFtYSJ9LCJleHAiOiIyMDIzLTEyLTE0VDA4OjQwOjI2KzA3OjAwIiwiaWF0IjoiMjAyMy0xMi0xM1QxNDo0MDoyNiswNzowMCIsImlkIjoiNjI4NTcyMjY5NzkxOCIsIm5iZiI6IjIwMjMtMTItMTNUMTQ6NDA6MjYrMDc6MDAifZMuLrRMEVB4gE_GOe8hLo6NrE_YGPn3NrIceJK_L0oSHtnlDzFlLKZXMf94I23k1ZGMlzl0VMhnRgKJrEt8eAg"
	filePath := "./image1.jpeg"
	formdataName := "file"

	responseBody, err := PostFileWithHeader(targetURL, tokenKey, tokenValue, filePath, formdataName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", string(responseBody))

}
