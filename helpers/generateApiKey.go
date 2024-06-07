package helpers

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/harsh082ip/URL-Shortener_Go/db"
	"go.mongodb.org/mongo-driver/bson"
)

func GenerateAPIKey(length int, ctx context.Context) (string, error) {
	var apikey string
	for {
		bytes := make([]byte, length)
		if _, err := rand.Read(bytes); err != nil {
			return "", err
		}
		apikey = hex.EncodeToString(bytes)
		collName := "ApiKeys"
		coll := db.OpenCollection(db.DBinstance(), collName)
		count, err := coll.CountDocuments(ctx, bson.M{"apikey": apikey})
		if err != nil {
			return "", fmt.Errorf("error checking for doc %v", err.Error())
		}
		if count > 0 {
			continue
		} else {
			break
		}
	}
	return apikey, nil
}
