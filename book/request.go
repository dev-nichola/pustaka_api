package book

import "encoding/json"

type BookRequest struct {
	Title       string      `json:"title" binding:"required,min=1,max=255"`
	Description string      `json:"description"`
	Price       json.Number `json:"price" binding:"required,numeric"`
	Rating      json.Number `json:"rating" binding:"omitempty,min=1,max=5"`
	Discount    json.Number `json:"discount" binding:"omitempty,min=0,max=100"`
}
