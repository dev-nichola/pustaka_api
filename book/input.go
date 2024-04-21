package book

import (
	"encoding/json"
)

type BookInput struct {
	Title    string      `json:"title" binding:"required"`
	Price    json.Number `json:"price" binding:"required,number"`
	Rating   int         `json:"rating" binding:"omitempty,min=1,max=5"`
	Discount int         `json:"discount" binding:"omitempty,min=0,max=100"`
}
