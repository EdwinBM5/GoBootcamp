package internal

import (
	"encoding/json"
	"fmt"
)

const (
	ErrProductNotFound = "Product not found"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (p *Product) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, Quantity: %d, CodeValue: %s, IsPublished: %t, Expiration: %s, Price: %f", p.ID, p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price)
}

func (p *Product) PickleProductJSON() (jsonData []byte, err error) {
	jsonData, err = json.Marshal(p)
	if err != nil {
		return
	}

	return
}
