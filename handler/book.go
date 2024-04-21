package handler

import (
	"fmt"
	"net/http"
	"pustaka_api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")

	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.FindById(int(id))

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()

	if err != nil {
		panic(err)
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := book.BookResponse{
			ID:       b.ID,
			Title:    b.Title,
			Price:    b.Price,
			Rating:   b.Rating,
			Discount: b.Discount,
		}

		booksResponse = append(booksResponse, bookResponse)

	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) PostBooksHandler(c *gin.Context) {
	var BookRequest book.BookRequest

	err := c.ShouldBindJSON(&BookRequest)

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

	book, err := h.bookService.Create(BookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

// func (h *bookHandler) RootHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"Bio":  "Calon Software Engginer",
// 		"Name": "Nichola Saputra",
// 	})
// }

// func (h *bookHandler) BookHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	title := c.Param("title")

// 	c.JSON(http.StatusOK, gin.H{
// 		"ID":    id,
// 		"Title": title,
// 	})
// }

// func (h *bookHandler) QueryHandler(c *gin.Context) {
// 	id := c.Query("title")
// 	price := c.Query("price")

// 	c.JSON(http.StatusOK, gin.H{
// 		"title": id,
// 		"price": price,
// 	})
// }

// func (h *bookHandler) HelloHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"name": "Hello World",
// 	})
// }
