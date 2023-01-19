package book

import "encoding/json"

type BookInput struct {
	Name  string      `json:"name" binding:"required"`
	Harga json.Number `json:"harga" binding:"required,number"`
	// Subtitel string `json:"sub_title"` //disini tak boleh salah sesuai dengan port nya
}
