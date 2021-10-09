package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserId          string             `json:"userId,omitempty" bson:"userId,omitempty"`
	Caption         string             `json:"caption,omitempty" bson:"caption,omitempty"`
	ImageURL        string             `json:"imageURL,omitempty" bson:"imageURL,omitempty"`
	PostedTimeStamp time.Time          `json:"postedTimeStamp,omitempty" bson:"postedTimeStamp,omitempty"`
}
