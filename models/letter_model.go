package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Letter struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Name    string             `json:"name,omitempty" validate:"required"`
	Content string             `json:"content,omitempty" validate:"required"`
	Title   string             `json:"title,omitempty" validate:"required"`
}
