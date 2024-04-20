package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	v2 := router.Group("/v2")

	v1.GET("/", rootHandler)
	v1.GET("/hello", helloHandler)
	v1.GET("/book/:id/title", bookHandler)
	v2.GET("/query", queryHandler)
	v2.POST("/books", postBooksHandler)

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
	Title      string      `json:"title" binding:"required"`
	Price      json.Number `json:"price" binding:"required,number"`
	BookAuthor string      `json:"book_author"`
}

func postBooksHandler(c *gin.Context) {
	var BookInput BookInput

	err := c.ShouldBindJSON(&BookInput)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error On Field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors":      errorMessages,
			"status code": http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":       BookInput.Title,
		"harga":       BookInput.Price,
		"book_author": BookInput.BookAuthor,
	})

}
