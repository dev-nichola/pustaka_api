package router

import (
	"pustaka_api/book"
	"pustaka_api/database"
	"pustaka_api/handler"

	"github.com/gin-gonic/gin"
)

func Router() {

	db, err := database.GetConnection()

	if err != nil {
		panic(err)
	}

	bookRepository := book.NewRepository(db)
	// bookFileRepository := book.NewFileRepository()
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")

	// v1.GET("/", bookHandler.RootHandler)
	// v1.GET("/hello", bookHandler.HelloHandler)
	// v1.GET("/book/:id/title", bookHandler.BookHandler)
	// v1.GET("/query", bookHandler.QueryHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)

	router.Run(":8000")
}
