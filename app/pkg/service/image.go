package service

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	model "github.com/SaRgEX/Diplom/Model"
)

type ImageService struct {
}

func NewImageService() *ImageService {
	return &ImageService{}
}

func (s *ImageService) UploadImage(file model.ImageInput) (string, error) {
	newFileName := fmt.Sprintf("%s.%s", RandomFileName(), file.Extension)
	dir := os.Getenv("IMAGE_DIR")

	fileOnDisk, err := os.Create(dir + "/" + newFileName)
	if err != nil {
		return "", err
	}

	defer fileOnDisk.Close()

	if _, err := fileOnDisk.Write(file.Image); err != nil {
		return "", err
	}
	return newFileName, nil
}

func RandomFileName() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return strings.ToLower(fmt.Sprintf("%v", entropy.Int63()))
}
