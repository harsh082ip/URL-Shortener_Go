package helpers

import (
	"context"
	"log"

	"github.com/harsh082ip/URL-Shortener_Go/db"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckIfDocExists(key, val, collName string, ctx context.Context) (bool, error) {

	coll := db.OpenCollection(db.DBinstance(), collName)
	filter := bson.M{key: val}
	res, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	log.Println(res)
	if res > 0 {
		return true, nil
	}
	return false, nil
}
