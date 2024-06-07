package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/URL-Shortener_Go/db"
	"github.com/harsh082ip/URL-Shortener_Go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID := c.Query("sessionID")
		if sessionID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Session ID is required",
				"error":  "Missing sessionID query parameter",
			})
			c.Abort()
			return
		}
		ctx, _ := context.WithTimeout(context.Background(), time.Second*120)

		// extract email of the sessionId
		collName := "SessionInfo"
		coll := db.OpenCollection(db.DBinstance(), collName)
		var sessionInfo models.SessionInfo
		err := coll.FindOne(ctx, bson.M{"sessionid": sessionID}).Decode(&sessionInfo)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status": "No user found for the sessionID",
					"error":  err.Error(),
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in finding user for the session ID",
				"error":  err.Error(),
			})
			c.Abort()
			return
		}

		rdb := db.RedisConnect()
		key := "session:" + sessionInfo.Email
		res, err := rdb.Exists(ctx, key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Error in checking for authorization status",
				"error":  err.Error(),
			})
			c.Abort()
			return
		}

		if res > 0 {
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "Unauthorized Access",
			"error":  "Session expired or doesn't exists",
		})
		c.Abort()
		// return
	}
}
