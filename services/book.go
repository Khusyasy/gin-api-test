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
	FindByID(string) (entities.Book, error)
	UpdateByID(string, entities.Book) (entities.Book, error)
	DeleteByID(string) error
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

func (service *bookService) FindByID(id string) (entities.Book, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entities.Book{}, err
	}

	var book entities.Book
	err = service.coll.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&book)
	if err != nil {
		return entities.Book{}, err
	}

	return book, nil
}

func (service *bookService) UpdateByID(id string, book entities.Book) (entities.Book, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entities.Book{}, err
	}

	_, err = service.coll.UpdateOne(context.Background(), bson.M{"_id": oid}, bson.M{"$set": book})
	if err != nil {
		return entities.Book{}, err
	}

	book.ID = oid
	return book, nil
}

func (service *bookService) DeleteByID(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = service.coll.DeleteOne(context.Background(), bson.M{"_id": oid})
	if err != nil {
		return err
	}

	return nil
}
