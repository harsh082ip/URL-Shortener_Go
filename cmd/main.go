package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/URL-Shortener_Go/routes"
)

const (
	WEBPORT = ":8002"
)

func main() {

	router := gin.Default()
	routes.AuthRoutes(router)

	if err := http.ListenAndServe(WEBPORT, router); err != nil {
		log.Fatal("Error Starting the server ", err.Error())
	}
}
