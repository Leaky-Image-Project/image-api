package service

import (
	"fmt"
	"leaky-image-project/chat-api/dto"
	"leaky-image-project/chat-api/entity"
)

type ImageService interface {
	Upload(i dto.ImageDTO) entity.Image
}

type imageService struct {
}

func NewImageService() ImageService {
	return &imageService{}
}

func (service *imageService) Upload(i dto.ImageDTO) entity.Image {
	file := i.ImgData
	fmt.Print(file.Filename)
	return entity.Image{}
}
