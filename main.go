package main

import (
	"context"
	"log"
	"os"

	"github.com/Khusyasy/gin-api-test/controllers"
	"github.com/Khusyasy/gin-api-test/services"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	dbname := os.Getenv("MONGODB_DBNAME")
	if dbname == "" {
		log.Fatal("You must set your 'MONGODB_DBNAME' environmental variable.")
	}
	db := client.Database(dbname)

	var (
		bookService services.BookService = services.NewBookService(db.Collection("books"))
		bookController controllers.BookController = controllers.NewBookController(bookService)
	)

	server := gin.Default()

	server.GET("/books", bookController.FindAll)
	server.POST("/books", bookController.Save)

	server.Run(":8080")
}
