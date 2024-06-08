package urlcontrollers

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/URL-Shortener_Go/db"
	"github.com/harsh082ip/URL-Shortener_Go/helpers"
	"github.com/harsh082ip/URL-Shortener_Go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Shorten(c *gin.Context) {

	var url models.UrlInfo
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "redirect url and email is required",
			"error":  err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	// Assuming url.RedirectURL is a string

	// Check if the URL starts with "https://"
	if strings.HasPrefix(url.RedirectURL, "https://") {
		// URL starts with "https://", it's already correct
	} else if strings.HasPrefix(url.RedirectURL, "http://") {
		// URL starts with "http://", it's already correct
	} else {
		// URL doesn't start with "https://" or "http://", so prepend "http://"
		url.RedirectURL = "http://" + url.RedirectURL
	}
	isUrlInvalid := helpers.IsValidURL(url.RedirectURL)
	if !isUrlInvalid {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "URL does not respond",
			"error":  "Entered Url doesn't work, pls enter correct url",
		})
		return
	}
	userExists, err := helpers.CheckIfDocExists("email", url.CreatedBy, "Users", ctx)
	if err != nil {
		log.Println("Error checking if user exists:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error checking user existence",
			"error":  err.Error(),
		})
		return
	}
	if !userExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Please send the correct email",
			"error":  "No user exists with this email",
		})
		return
	}
	shortID, err := helpers.CreateShortID(url.RedirectURL, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error in creating shortID",
			"error":  err.Error(),
		})
		return
	}

	url.ShortID = shortID
	// url.CreatedAt = time.Now()
	url.ID = primitive.NewObjectID()

	collName := "UrlInfo"
	coll := db.OpenCollection(db.DBinstance(), collName)
	// log.Println(coll)

	// _, err = coll.InsertOne(ctx, url)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"status":  "Error in setting the short id",
	// 		"error":   err.Error(),
	// 		"UrlInfo": url,
	// 	})
	// 	return
	// }

	// first we will check if the same user has previously created the shorturl of the same
	filter := bson.M{"createdby": url.CreatedBy, "redirecturl": url.RedirectURL, "requesttype": "non_custom"}
	count, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error in checking existing data",
			"error":  err.Error(),
		})
		return
	}

	if count == 0 {
		url.RequestType = "non_custom"
		url.CreatedAt = time.Now()
		url.UpdatedAt = url.CreatedAt
		_, err = coll.InsertOne(ctx, url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Error in setting the short id",
				"error":   err.Error(),
				"UrlInfo": url,
			})
			return
		}
	} else {
		update := bson.M{"$set": bson.M{"shortid": url.ShortID, "updatedAt": time.Now()}}

		opts := options.Update().SetUpsert(false) // SetUpsert(false) for update only

		_, err := coll.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error in updating existing shortid",
				"error":  err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})

}
