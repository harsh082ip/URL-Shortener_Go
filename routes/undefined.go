package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UndefinedRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", gin.H{
			"title": "Page Not Found",
		})
	})
}
