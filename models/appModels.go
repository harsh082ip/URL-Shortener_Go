package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApiKey struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email  string             `json:"email" binding:"required,email"`
	ApiKey string             `json:"apikey" binding:"required,min=32"`
}

type SessionInfo struct {
	Email     string    `json:"email" binding:"required,email"`
	SessionID string    `json:"sessionid" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedNow"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `json:"name" binding:"required"`
	Email    string             `json:"email" binding:"required,email"`
	Password string             `json:"password" binding:"required,min=6"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// type visithistory struct {
// 	// ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
// 	Timestamp []time.Time `bson:"timestamp"`
// }

type UrlInfo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ShortID     string             `bson:"shortid"`
	RedirectURL string             `bson:"redirecturl" binding:"required"`
	// VisitHistory visithistory       `bson:"visithistory,omitempty"`
	CreatedBy string    `bson:"createdby" binding:"required"`
	CreatedAt time.Time `bson:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty"`
	// VisitCount   int                `bson:"visitcount"`
}

type VisitsInfoCache struct {
	VisitCount int         `json:"visitcount"`
	Timestamps []time.Time `json:"timestamps"`
	SessionID  string      `json:"sessionid"`
}
