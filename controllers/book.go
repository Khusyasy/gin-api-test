package controllers

import (
	"github.com/Khusyasy/gin-api-test/entities"
	"github.com/Khusyasy/gin-api-test/services"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	Save(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type controller struct {
	service services.BookService
}

func NewBookController(service services.BookService) BookController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) {
	var book entities.Book
	ctx.BindJSON(&book)
	c.service.Save(book)
	ctx.JSON(200, book)
}

func (c *controller) FindAll(ctx *gin.Context) {
	var books []entities.Book = c.service.FindAll()
	ctx.JSON(200, books)
}
