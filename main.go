package main

import (
	"github.com/Khusyasy/gin-api-test/controllers"
	"github.com/Khusyasy/gin-api-test/services"
	"github.com/gin-gonic/gin"
)

var (
	bookService services.BookService = services.NewBookService()
	bookController controllers.BookController = controllers.NewBookController(bookService)
)

func main() {
	server := gin.Default()

	server.GET("/books", bookController.FindAll)
	server.POST("/books", bookController.Save)

	server.Run(":8080")
}
