package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	// log.Fatal(MONGODB_URI)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*15)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Client", client)

	return client

}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("Url-Shortner").Collection(collectionName)
	return collection
}
