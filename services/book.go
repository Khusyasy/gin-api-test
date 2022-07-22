package services

import (
	"context"

	"github.com/Khusyasy/gin-api-test/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookService interface {
	Save(entities.Book) (entities.Book, error)
	FindAll() ([]entities.Book, error)
}

type bookService struct {
	coll *mongo.Collection
}

func NewBookService(coll *mongo.Collection) BookService {
	return &bookService{
		coll: coll,
	}
}

func (service *bookService) Save(book entities.Book) (entities.Book, error) {
	book.ID = primitive.NewObjectID()

	_, err := service.coll.InsertOne(context.Background(), book)
	if err != nil {
		return entities.Book{}, err
	}

	return book, nil
}

func (service *bookService) FindAll() ([]entities.Book, error) {
	cur, err := service.coll.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var books []entities.Book
	for cur.Next(context.Background()) {
		var book entities.Book
		err := cur.Decode(&book)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.Background())
	return books, nil
}
