package internal

import (
	"encoding/json"
	"fmt"
)

type StorageProduct struct {
	Products []Product `json:"products"`
}

func (s *StorageProduct) AddProduct(p Product) {
	s.Products = append(s.Products, p)
}

func (s *StorageProduct) GetProducts() []Product {
	return s.Products
}

func (s *StorageProduct) String() {
	for _, product := range s.Products {
		fmt.Println(product.String())
	}
}

func (s *StorageProduct) PickleProductJSON() (jsonData []byte, err error) {
	jsonData, err = json.Marshal(s.Products)
	if err != nil {
		return
	}

	return
}

func (s *StorageProduct) GetProductByID(id int) (Product, error) {
	for _, product := range s.Products {
		if product.ID == id {
			return product, nil
		}
	}

	return Product{}, fmt.Errorf(ErrProductNotFound)
}

func (s StorageProduct) GetProductByPrice(price float64) (*StorageProduct, error) {
	var products StorageProduct
	for _, product := range s.Products {
		if product.Price > price {
			products.Products = append(products.Products, product)
		}
	}

	if len(products.Products) == 0 {
		return nil, fmt.Errorf(ErrProductNotFound)
	}

	return &products, nil
}
