package services

import "github.com/Khusyasy/gin-api-test/entities"

type BookService interface {
	Save(entities.Book) entities.Book
	FindAll() []entities.Book
}

type bookService struct {
	books []entities.Book
}

func NewBookService() BookService {
	return &bookService{}
}

func (service *bookService) Save(book entities.Book) entities.Book {
	service.books = append(service.books, book)
	return book
}

func (service *bookService) FindAll() []entities.Book {
	return service.books
}
