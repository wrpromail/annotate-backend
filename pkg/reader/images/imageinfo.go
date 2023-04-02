package images

import (
	"bufio"
	"github.com/google/uuid"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

type Photo struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Size   int64  `json:"size"`
}

type Data struct {
	Photos []Photo `json:"photos,omitempty"`
}

type Response struct {
	Success bool   `json:"success"`
	Data    Data   `json:"data"`
	Error   string `json:"error,omitempty"`
}

func getJpgImageDimensions(imgPath string) (int, int, error) {
	inputFile, err := os.Open(imgPath)
	if err != nil {
		return 0, 0, err
	}
	reader := bufio.NewReader(inputFile)
	config, _, err := image.DecodeConfig(reader)
	if err != nil {
		return 0, 0, err
	}
	return config.Width, config.Height, nil
}

func scanDir(root string) ([]Photo, error) {
	var photos []Photo
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".jpg" {
			return nil
		}

		width, height, _ := getJpgImageDimensions(path)
		if err != nil {
			// ?
		}

		photo := Photo{
			ID:     uuid.New().String(),
			Name:   info.Name(),
			Path:   path,
			Width:  width,
			Height: height,
			Size:   info.Size() / 1024,
		}
		photos = append(photos, photo)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return photos, nil
}
