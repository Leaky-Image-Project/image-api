package controller

import (
	"fmt"
	"leaky-image-project/chat-api/dto"
	"leaky-image-project/chat-api/entity"
	"leaky-image-project/chat-api/helper"
	"leaky-image-project/chat-api/service"
	"net/http"
	"strconv"

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

	authHeader := ctx.GetHeader("Authorization")
	_, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	res := c.imageService.Upload(imageUploadDTO)
	if (res == entity.ImageInfo{}) {
		res := helper.BuildErrorResponse("Internal error", "File upload error", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		response := helper.BuildResponse(true, "OK", res)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *imageController) DownloadImage(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	_, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	fmt.Println(id)
}
