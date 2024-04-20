package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/book/:id/title", bookHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", postBooksHandler)

	router.Run(":8000")
}

func bookHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"ID":    id,
		"Title": title,
	})
}

func queryHandler(c *gin.Context) {
	id := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": id,
		"price": price,
	})
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Bio":  "Calon Software Engginer",
		"Name": "Nichola Saputra",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Hello World",
	})
}

type BookInput struct {
	Title      string `json:"title" binding:"required"`
	Price      int    `json:"price" binding:"required,number"`
	BookAuthor string `json:"book_author"`
}

func postBooksHandler(c *gin.Context) {
	var BookInput BookInput

	err := c.ShouldBindJSON(&BookInput)

	if err != nil {

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error On Field %s, condition %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)

			return
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"title":       BookInput.Title,
		"harga":       BookInput.Price,
		"book_author": BookInput.BookAuthor,
	})

}
