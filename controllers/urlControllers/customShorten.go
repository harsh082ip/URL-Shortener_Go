package urlcontrollers

import (
	"context"
	"encoding/json"
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

func CustomShorten(c *gin.Context) {

	var url models.UrlInfo
	if err := json.NewDecoder(c.Request.Body).Decode(&url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error in Decoding Json Body",
			"error":  err.Error(),
		})
		return
	}
	log.Println(url.ShortID)
	log.Println(url.RedirectURL)
	log.Println(url.CreatedBy)
	if url.ShortID == "" || url.RedirectURL == "" || url.CreatedBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Details Missing",
			"error":  "ShortID, RedirectUrl and CreatedBy is required to proceed",
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	if len(url.ShortID) >= 6 && len(url.ShortID) <= 8 {
		shortIdExists, err := helpers.CheckIfDocExists("shortid", url.ShortID, "UrlInfo", ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in checking for shortID in db",
				"error":  err.Error(),
			})
			return
		}
		if shortIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ShortID already exists",
				"error":  "Please request with another shortID",
			})
			return
		}

		// Check if the URL starts with "https://"
		if strings.HasPrefix(url.RedirectURL, "https://") {
			// URL starts with "https://", it's already correct
		} else if strings.HasPrefix(url.RedirectURL, "http://") {
			// URL starts with "http://", it's already correct
		} else {
			// URL doesn't start with "https://" or "http://", so prepend "http://"
			url.RedirectURL = "http://" + url.RedirectURL
		}

		isUrlValid := helpers.IsValidURL(url.RedirectURL)
		if !isUrlValid {
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

		url.ID = primitive.NewObjectID()
		collName := "UrlInfo"
		coll := db.OpenCollection(db.DBinstance(), collName)

		// first we will check if the same user has previously created the shorturl of the same
		filter := bson.M{"createdby": url.CreatedBy, "redirecturl": url.RedirectURL, "requesttype": "custom"}
		count, err := coll.CountDocuments(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error in checking existing data",
				"error":  err.Error(),
			})
			return
		}

		if count == 0 {
			url.RequestType = "custom"
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
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"status": "length not correct",
		"error":  "Length should be between 6 and 8 characters",
	})
	// return
}
