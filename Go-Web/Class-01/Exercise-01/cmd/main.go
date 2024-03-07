package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	/*
		We are going to create a web application with the net/http package native to go,
		that has a /ping endpoint that when pasted responds a text that says "pong".

		The endpoint should be a GET method
		The response of "pong" must be sent as text, NOT as JSON.
	*/
	router := chi.NewRouter()
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error starting server: ", err.Error())
	}
}
