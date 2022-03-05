package main

import (
	"leaky-image-project/chat-api/controller"
	"leaky-image-project/chat-api/middleware"
	"leaky-image-project/chat-api/service"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

var (
	jwtService      service.JWTService         = service.NewJWTService()
	authService     service.AuthService        = service.NewAuthService()
	imageService    service.ImageService       = service.NewImageService()
	authController  controller.AuthController  = controller.NewAuthController(authService, jwtService)
	imageController controller.ImageController = controller.NewImageController(imageService, jwtService)
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", authController.Login)
	}

	imageRoute := r.Group("/image", middleware.AuthorizeJWT(jwtService))
	{
		imageRoute.POST("/upload", imageController.UploadImage)
		imageRoute.GET("/:id", imageController.DownloadImage)
	}

	r.Run(":3030")
}
