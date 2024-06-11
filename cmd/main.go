package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/URL-Shortener_Go/consts"
	"github.com/harsh082ip/URL-Shortener_Go/routes"
)

func main() {
	router := gin.Default()

	// Load HTML templates from the templates directory
	// The LoadHTMLGlob function tells Gin where to find your HTML templates.
	router.LoadHTMLGlob("templates/*")

	// Defined routes
	routes.AuthRoutes(router)
	routes.UrlRoutes(router)
	routes.MiscRoutes(router)

	// for un-defined routes
	routes.UndefinedRoutes(router)
	// Serve static files
	// router.Static("/tempelates", "./tempelates")

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Welcome to url shortener service :)",
		})
	})

	// Start the server
	if err := http.ListenAndServe(consts.WEBPORT, router); err != nil {
		log.Fatal("Error Starting the server ", err.Error())
	}
}
