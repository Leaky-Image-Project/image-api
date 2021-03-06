package controller

import (
	"leaky-image-project/image-api/dto"
	"leaky-image-project/image-api/helper"
	"leaky-image-project/image-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var credentials dto.Credentials
	errDTO := ctx.ShouldBind(&credentials)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	isAuthenticated := c.authService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		generatedToken := c.jwtService.GenerateToken(credentials.Username)
		ctx.SetCookie("token", generatedToken, 3600, "/", "localhost", false, true)
		response := helper.BuildResponse(true, "OK!", generatedToken)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Logout(ctx *gin.Context) {
	// a fake logout just delete the cookie
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	response := helper.BuildResponse(true, "OK!", "Token has been deleted")
	ctx.JSON(http.StatusOK, response)
}
