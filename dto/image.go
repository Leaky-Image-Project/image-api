package dto

import "mime/multipart"

type ImageDTO struct {
	ImgData *multipart.FileHeader `form:"img_data" binding:"required"`
}
