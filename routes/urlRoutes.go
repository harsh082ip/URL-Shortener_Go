package routes

import (
	"github.com/gin-gonic/gin"
	urlcontrollers "github.com/harsh082ip/URL-Shortener_Go/controllers/urlControllers"
	"github.com/harsh082ip/URL-Shortener_Go/middleware"
)

func UrlRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/url/shorten", middleware.AuthMiddleware(), urlcontrollers.Shorten)
}
