package service

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"leaky-image-project/chat-api/dto"
	"leaky-image-project/chat-api/entity"
	"leaky-image-project/chat-api/helper"
	"os"
)

type ImageService interface {
	Upload(i dto.ImageUploadDTO) entity.ImageInfo
}

type imageService struct {
}

func NewImageService() ImageService {
	return &imageService{}
}

func (service *imageService) Upload(i dto.ImageUploadDTO) entity.ImageInfo {
	fileHeader := i.ImgData
	fmt.Print(fileHeader.Filename)

	file, err := fileHeader.Open()
	if err != nil {
		// TODO: check file IO error
		return entity.ImageInfo{}
	}
	defer file.Close()

	bufFile := bufio.NewReader(file)

	image, imageType, err := image.Decode(bufFile)
	if err != nil {
		// TODO: check decode err
		return entity.ImageInfo{}
	}

	if !helper.HasType(imageType) {
		// TODO: check supported type
		return entity.ImageInfo{}
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		// TODO: check moving position
		return entity.ImageInfo{}
	}

	md5Hash := md5.New()
	// bufFile := bufio.NewReader(file)
	_, err = io.Copy(md5Hash, bufFile)

	if err != nil {
		//TODO: md5 encoding
		return entity.ImageInfo{}
	}

	fileMd5Fx := md5Hash.Sum(nil)
	fileMd5 := fmt.Sprintf("%x", fileMd5Fx)

	dirPath := helper.JoinPath(fileMd5) + "/"
	filePath := dirPath + fileMd5

	dirInfo, err := os.Stat(dirPath)
	if err != nil {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			// TODO: file path err
			return entity.ImageInfo{}
		}
	} else {
		if !dirInfo.IsDir() {
			err = os.MkdirAll(dirPath, 0755)
			if err != nil {
				// TODO: file path err
				return entity.ImageInfo{}
			}
		}
	}

	_, err = os.Stat(filePath)
	if err != nil {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			// TODO: file err
			return entity.ImageInfo{}
		}
		defer file.Close()

		if imageType == helper.PNG {
			err = png.Encode(file, image)
		} else if imageType == helper.JPG || imageType == helper.JPEG {
			err = jpeg.Encode(file, image, nil)
		}

		if err != nil {
			// TODO: encoding error
			return entity.ImageInfo{}
		}
	}

	return entity.ImageInfo{
		Id:   fileMd5,
		Size: fileHeader.Size,
		Mime: imageType,
	}
}
