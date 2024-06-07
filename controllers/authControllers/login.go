package authcontrollers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/URL-Shortener_Go/consts"
	"github.com/harsh082ip/URL-Shortener_Go/db"
	"github.com/harsh082ip/URL-Shortener_Go/helpers"
	"github.com/harsh082ip/URL-Shortener_Go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(c *gin.Context) {

	var jsonData models.LoginUser
	// var apikey models.ApiKey
	// Bind and validate JSON
	var sessionInfo models.SessionInfo
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		// Return a bad request response if there's an error in binding/validation
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error in request body",
			"error":  err.Error(),
		})
		return
	}

	hashPass := jsonData.Password
	collName := "Users"
	coll := db.OpenCollection(db.DBinstance(), collName)
	rdb := db.RedisConnect()
	ctx, _ := context.WithTimeout(context.Background(), time.Second*15)

	err := coll.FindOne(ctx, bson.M{"email": jsonData.Email}).Decode(&jsonData)
	if err != nil {
		// Handle document not found error
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Error in Document",
				"error":  "No User found with the given details",
			})
			return
		}
		// Handle internal server error while searching for the user
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Internal Server Error",
			"error":  "Error is searching for user",
		})
		return
	}

	err = helpers.ComparePassword(jsonData.Password, hashPass)
	if err != nil {
		// Return unauthorized access response if passwords do not match
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "Unauthorized Access",
			"error":  "Password Mismatch",
		})
		return
	}

	sessionInfo.SessionID, err = helpers.CreateSeessionID(consts.SessionIDlength, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "SessionID generation error",
			"error":  err.Error(),
		})
		return
	}

	sessionInfo.Email = jsonData.Email

	// ----------------- SET to REDIS ----------------------------
	key := "session:" + sessionInfo.SessionID
	jsonSession, err := json.Marshal(sessionInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error in Marshalling struct",
			"error":  err.Error(),
		})
		return
	}
	_, err = rdb.Set(ctx, key, jsonSession, consts.SessionTTL).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error in Setting SessionID to DB",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "User SignUp Successful",
		"sessionInfo": sessionInfo,
	})
}
