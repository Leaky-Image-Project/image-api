package dto

import "mime/multipart"

type ImageUploadDTO struct {
	ImgData *multipart.FileHeader `form:"img_data" binding:"required"`
}

type ImageDownloadDTO struct {
	Id string `uri:"id" binding:"required"`
}
