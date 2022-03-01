package main

import (
	"leaky-image-project/chat-api/controller"
	"leaky-image-project/chat-api/middleware"
	"leaky-image-project/chat-api/service"

	"github.com/gin-gonic/gin"
)

var (
	jwtService      service.JWTService         = service.NewJWTService()
	authService     service.AuthService        = service.NewAuthService()
	authController  controller.AuthController  = controller.NewAuthController(authService, jwtService)
	imageController controller.ImageController = controller.NewImageController(jwtService)
)

func main() {
	r := gin.Default()

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", authController.Login)
	}

	imageRoute := r.Group("/image", middleware.AuthorizeJWT(jwtService))
	{
		imageRoute.POST("/upload", imageController.UploadImage)
		imageRoute.GET("/:id", imageController.DownloadImage)
	}

	r.Run(":3000")
}
