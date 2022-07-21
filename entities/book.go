package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID   primitive.ObjectID `bson:"_id"`
	Title string `bson:"title"`
	Author string `bson:"author"`
	Year string `bson:"year"`
}
