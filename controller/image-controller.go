package controller

import (
	"io"
	"leaky-image-project/image-api/dto"
	"leaky-image-project/image-api/helper"
	"leaky-image-project/image-api/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ImageController interface {
	UploadImage(ctx *gin.Context)
	DownloadImage(ctx *gin.Context)
}

type imageController struct {
	imageService service.ImageService
	jwtService   service.JWTService
}

func NewImageController(imageService service.ImageService, jwtService service.JWTService) ImageController {
	return &imageController{
		jwtService:   jwtService,
		imageService: imageService,
	}
}

func (c *imageController) UploadImage(ctx *gin.Context) {
	var imageUploadDTO dto.ImageUploadDTO
	errDTO := ctx.ShouldBind(&imageUploadDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// authHeader := ctx.GetHeader("Authorization")
	authHeader, _ := ctx.Cookie("token")
	_, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	res, err := c.imageService.Upload(imageUploadDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Internal error", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		response := helper.BuildResponse(true, "OK", res)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *imageController) DownloadImage(ctx *gin.Context) {
	var imageDownloadDTO dto.ImageDownloadDTO
	errDto := ctx.ShouldBindUri(&imageDownloadDTO)
	if errDto != nil {
		res := helper.BuildErrorResponse("No param id was found", errDto.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// authHeader := ctx.GetHeader("Authorization")
	authHeader, _ := ctx.Cookie("token")
	_, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}

	filePath := helper.UrlParse(imageDownloadDTO.Id)
	file, err := os.Open(filePath)
	if err != nil {
		res := helper.BuildErrorResponse("File not exist", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	io.Copy(ctx.Writer, file)
}
