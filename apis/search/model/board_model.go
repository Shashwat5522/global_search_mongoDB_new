package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Board struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}