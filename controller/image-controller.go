package controller

import (
	"fmt"
	"leaky-image-project/chat-api/dto"
	"leaky-image-project/chat-api/helper"
	"leaky-image-project/chat-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageController interface {
	UploadImage(ctx *gin.Context)
	DownloadImage(ctx *gin.Context)
}

type imageController struct {
	jwtService service.JWTService
}

func NewImageController(jwtService service.JWTService) ImageController {
	return &imageController{
		jwtService: jwtService,
	}
}

func (c *imageController) UploadImage(ctx *gin.Context) {
	var imageDTO dto.ImageDTO
	errDTO := ctx.ShouldBind(&imageDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	file := imageDTO.ImgData
	fmt.Print(file.Filename)
	// authHeader := ctx.GetHeader("Authorization")
	// _, errToken := c.jwtService.ValidateToken(authHeader)
	// if errToken != nil {
	// 	panic(errToken.Error())
	// }

}

func (c *imageController) DownloadImage(ctx *gin.Context) {

}
