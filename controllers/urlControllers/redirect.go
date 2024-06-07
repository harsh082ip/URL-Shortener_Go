package urlcontrollers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/URL-Shortener_Go/db"
	"github.com/harsh082ip/URL-Shortener_Go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RedirectUrlS(c *gin.Context) {
	log.Println("In Redirect Url...")
	shortID := c.Param("shortid")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	var url models.UrlInfo

	// Check Redis if info is available
	rdb := db.RedisConnect()
	if rdb == nil {
		log.Println("Error: Redis connection failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis connection failed"})
		return
	}

	key := "shortid:" + shortID
	res, err := rdb.Get(ctx, key).Result()
	if err == nil && res != "" {
		log.Println("Nillll")
		err := json.Unmarshal([]byte(res), &url)
		if err != nil {
			log.Println("Error in Unmarshaling Redis data", err.Error())
		} else {
			if url.RedirectURL != "" {
				log.Println("redis hehehhe", url.RedirectURL)
				c.Redirect(http.StatusMovedPermanently, url.RedirectURL)
				log.Println("Redirected to:", url.RedirectURL)
				return // Early return to avoid further processing
			}
		}
	}

	// If not present in Redis, check Mongo and then set it to Redis
	collName := "UrlInfo"
	coll := db.OpenCollection(db.DBinstance(), collName)

	err = coll.FindOne(ctx, bson.M{"shortid": shortID}).Decode(&url)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// shortid does not exist, serve 404 page
			log.Println("Short ID not found, serving 404 page")
			c.HTML(http.StatusNotFound, "404.html", gin.H{
				"title": "Page Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error in searching for shortid",
			"error":  err.Error(),
		})
		return
	}

	// ShortID found -- then set to Redis
	jsonData, err := json.Marshal(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error in marshaling the body from mongo",
			"error":  err.Error(),
		})
		return
	}
	res, err = rdb.Set(ctx, key, jsonData, time.Second*120).Result()
	if err != nil {
		log.Println("Error in Setting UrlInfo to Redis", err.Error())
	}
	fmt.Println(res)
	log.Println("Redirecting to URL mongo", url.RedirectURL)
	c.Redirect(http.StatusMovedPermanently, url.RedirectURL)
	log.Println("Redirected to:", url.RedirectURL)
}
