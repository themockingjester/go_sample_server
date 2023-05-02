package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	// Bind the server HTTP endpoints.
	r := chi.NewRouter()

	r.Post("/postRequest/{source}", PostRequest)
	r.HandleFunc("/about", AboutPage)
	r.Get("/getRequest", GetRequest)
	err := http.ListenAndServe(":9098", r)

	// These below print statements are not working as expected
	if err == nil {
		fmt.Println("Successfully started server")
	} else {
		fmt.Println("An Error occoured  %v", err)
	}
}
