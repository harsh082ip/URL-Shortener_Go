package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ApiKey struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email  string             `json:"email" binding:"required,email"`
	ApiKey string             `json:"apikey" binding:"required,min=32"`
}

type SessionInfo struct {
	Email     string `json:"email" binding:"required,email"`
	SessionID string `json:"sessionid" binding:"required"`
}
