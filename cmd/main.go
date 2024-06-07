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
	routes.AuthRoutes(router)
	routes.UrlRoutes(router)

	if err := http.ListenAndServe(consts.WEBPORT, router); err != nil {
		log.Fatal("Error Starting the server ", err.Error())
	}
}
