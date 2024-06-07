package routes

import (
	"github.com/gin-gonic/gin"
	authcontrollers "github.com/harsh082ip/URL-Shortener_Go/controllers/authControllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/auth/signup", authcontrollers.SignUp)
	incomingRoutes.POST("/auth/login", authcontrollers.Login)
}
