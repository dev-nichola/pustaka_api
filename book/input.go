package book

import (
	"encoding/json"
)

type BookInput struct {
	Title      string      `json:"title" binding:"required"`
	Price      json.Number `json:"price" binding:"required,number"`
	BookAuthor string      `json:"book_author"`
}
