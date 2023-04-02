package images

import (
	"bufio"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"testing"
)

func TestReadImageSize(t *testing.T) {
	img := "D:\\react-basic-data\\src\\miru\\Fsl_5GCacAETGR4.jpg"
	inputFile, err := os.Open(img)
	if err != nil {
		t.Error(err)
		return
	}

	reader := bufio.NewReader(inputFile)
	config, _, _ := image.DecodeConfig(reader)
	t.Log(fmt.Sprintf("width: %d, height: %d", config.Width, config.Height))

}

func TestScanFolder(t *testing.T) {
	rst, err := scanDir("D:\\react-basic-data\\src\\miru")
	if err != nil {
		t.Error(err)
		return
	}
	var result = Response{
		Error:   "",
		Success: true,
		Data:    Data{rst},
	}
	rb, _ := json.Marshal(result)
	t.Log(string(rb))
}
