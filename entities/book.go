package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title  string `bson:"title" json:"title"`
	Author string `bson:"author" json:"author"`
	Year   string `bson:"year" json:"year"`
}
