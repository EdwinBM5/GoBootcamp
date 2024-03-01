package internal

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Category: %s, Descripcion: %s\n", product.ID, product.Name, product.Price, product.Category, product.Description)
	}
}

func (p Product) Save(products []Product) {
	Products = append(products, p)
	fmt.Printf("\nThe product %s has been saved successfully.\n\n", p.Name)
}

func GetById(productId int) {
	var productFound bool
	for _, product := range Products {
		if product.ID == productId {
			fmt.Println("\nProduct find with the following information...")
			fmt.Printf("ID: %d, Name: %s, Price: %.2f, Category: %s, Descripcion: %s\n", product.ID, product.Name, product.Price, product.Category, product.Description)
			productFound = true
			break
		}
	}
	if !productFound {
		fmt.Printf("\nProduct with (ID: %d) not found...\n", productId)
	}

}

var Products []Product = []Product{
	{ID: 1, Name: "Milk", Description: "Milk from a cow", Price: 1.99, Category: "Food"},
	{ID: 2, Name: "Bread", Description: "Freshly baked bread", Price: 2.50, Category: "Food"},
	{ID: 3, Name: "Apple", Description: "Crisp, juicy apples", Price: 0.99, Category: "Food"},
	{ID: 4, Name: "Water Bottle", Description: "500ml bottled water", Price: 0.89, Category: "Beverages"},
	{ID: 5, Name: "Chocolate Bar", Description: "Rich, dark chocolate bar", Price: 1.49, Category: "Snacks"},
	{ID: 6, Name: "Notebook", Description: "100-page blank notebook", Price: 3.99, Category: "Stationery"},
	{ID: 7, Name: "Pen", Description: "Blue ballpoint pen", Price: 0.49, Category: "Stationery"},
	{ID: 8, Name: "Cheese", Description: "Aged cheddar cheese", Price: 5.99, Category: "Food"},
	{ID: 9, Name: "Coffee Beans", Description: "Premium Arabica coffee beans", Price: 14.99, Category: "Beverages"},
	{ID: 10, Name: "Tea", Description: "Green tea leaves", Price: 4.99, Category: "Beverages"},
}
