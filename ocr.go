package mapdf

import (
	"github.com/otiai10/gosseract"
)

func ReadContent(filepath string) (content string, err error) {
	client := gosseract.NewClient()
	defer client.Close()
	err = client.SetImage(filepath)
	if err != nil {
		return
	}
	err = client.SetLanguage("eng")
	if err != nil {
		return
	}
	content, err = client.Text()
	return

}
