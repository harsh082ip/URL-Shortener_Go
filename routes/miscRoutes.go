package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiscRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Everything is good here 👀",
		})
	})
}
