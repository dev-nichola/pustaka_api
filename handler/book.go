package handler

import (
	"fmt"
	"net/http"
	"pustaka_api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Bio":  "Calon Software Engginer",
		"Name": "Nichola Saputra",
	})
}

func BookHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"ID":    id,
		"Title": title,
	})
}

func QueryHandler(c *gin.Context) {
	id := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": id,
		"price": price,
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Hello World",
	})
}

func PostBooksHandler(c *gin.Context) {
	var BookInput book.BookInput

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
