package helpers

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/harsh082ip/URL-Shortener_Go/consts"
	"github.com/harsh082ip/URL-Shortener_Go/db"
	"github.com/spaolacci/murmur3"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateShortID(url string, ctx context.Context) (string, error) {
	for {
		salt, err := GenerateRandomSalt(20)
		if err != nil {
			return "", fmt.Errorf("error in creating short id: salt error, %v", err)
		}
		input := url + salt

		// Create a new 128-bit MurmurHash3 hasher
		hash128 := murmur3.New128()

		// Write the input string to the hasher
		hash128.Write([]byte(input))

		// Compute the 128-bit hash value
		hashValue128 := hash128.Sum(nil)

		// Encode the hash value using hexadecimal encoding
		shortID := hex.EncodeToString(hashValue128)

		collName := "UrlInfo"
		coll := db.OpenCollection(db.DBinstance(), collName)

		// check if shortid already present
		res, err := coll.CountDocuments(ctx, bson.M{"shortid": shortID})
		if err != nil {
			return "", err
		}

		if res > 0 {
			continue
		} else {
			return shortID[:6], nil
		}
	}
}

func GenerateRandomSalt(length int) (string, error) {
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(consts.LetterBytes))))
		if err != nil {
			return "", err
		}
		result[i] = consts.LetterBytes[num.Int64()]
	}

	salt := string(result)
	return salt, nil
}
