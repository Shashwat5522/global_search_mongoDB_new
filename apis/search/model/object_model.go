package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Object struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	BoardID     Board              `bson:"boardID,omitempty"`
	Visible     string             `bson:"visible,omitempty"`
	Tags        []string           `bson:"tags,omitempty"`
	Description string             `bson:"description,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Type        string             `bson:"type,omitempty"`
}
