package main

import "class/bases/five/product/internal"

func main() {

	product := internal.Product{}
	internal.Product.GetAll(product)

	product.ID = 11
	product.Name = "Cookies"
	product.Description = "Cookies made from flour, sugar, and butter"
	product.Price = 2.99
	product.Category = "Food"

	internal.Product.Save(product, internal.Products)
	internal.Product.GetAll(product)

	internal.GetById(15)
}
