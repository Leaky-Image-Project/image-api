package main

import (
	"leaky-image-project/image-api/controller"
	"leaky-image-project/image-api/middleware"
	"leaky-image-project/image-api/service"

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
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/logout", authController.Logout)
	}

	imageRoute := r.Group("/image", middleware.AuthorizeJWT(jwtService))
	{
		imageRoute.POST("/upload", imageController.UploadImage)
		imageRoute.GET("/:id", imageController.DownloadImage)
	}

	r.Run(":3030")
}
