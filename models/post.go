package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post represents model for post
type Post struct {
	ID      primitive.ObjectID
	Title   string
	Article string
}
