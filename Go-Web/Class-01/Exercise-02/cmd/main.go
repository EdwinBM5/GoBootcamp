package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

func main() {
	/*
		Let's create an endpoint called /greetings. With a small structure with name and surname that when pasted
		should respond in text "Hello + name + surname".

		The endpoint should be a POST method
		The JSON package should be used to solve the exercise.
		The response should follow this structure: "Hello Andrea Rivas".

		The structure should look like this:
		{
		   	"firstName": "Andrea",
			"lastName": "Rivas"
		}
	*/
	router := chi.NewRouter()
	router.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")

		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error: decoding request body"))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hello %s %s", user.FirstName, user.LastName)))
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error starting server: ", err.Error())
	}
}
