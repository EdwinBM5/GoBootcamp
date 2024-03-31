package main

import (
	"class/web/three/supermarket/internal"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var products internal.StorageProduct

const (
	ErrProductNotLoad = "Error: products cannot be loaded"
	ErrorInvalidID    = "Error: invalid product ID"
	ErrorInvalidPrice = "Error: invalid entry please verify the product price"
)

func init() {
	productsContent, err := internal.LoadProductJSON("products.json")
	if err != nil {
		fmt.Println(err)
	}

	products, err = internal.DumpProductJSON(productsContent, products)
	if err != nil {
		fmt.Println(err)
	}

	// products.String()
	fmt.Printf("Loading products from the supermarket, having a total of %d...\n", len(products.Products))
}

func handlerProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	productsJSON, err := internal.JSON(&products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ErrProductNotLoad))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(productsJSON)
}

func handlerProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := chi.URLParam(r, "id")

	productId, err := strconv.ParseFloat(id, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(ErrorInvalidID))
	}

	product, err := products.GetProductByID(int(productId))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	}

	productJSON, err := internal.JSON(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ErrProductNotLoad))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(productJSON)
}

func handlerProductSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	price := r.URL.Query().Get("price")

	productPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(ErrorInvalidPrice))
		return
	}

	product, err := products.GetProductByPrice(productPrice)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	productJSON, err := internal.JSON(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ErrProductNotLoad))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(productJSON)
}

func main() {
	router := chi.NewRouter()
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	router.Route("/products", func(r chi.Router) {
		r.Get("/", handlerProducts)
		r.Get("/{id}", handlerProductByID)
		r.Get("/search", handlerProductSearch)
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error starting server: ", err.Error())
	}

}
