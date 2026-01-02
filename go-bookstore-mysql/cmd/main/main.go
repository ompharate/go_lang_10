package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ompharate/go-bookstore-mysql/pkg/config"
	"github.com/ompharate/go-bookstore-mysql/pkg/routes"
)

func main() {
	config.Connect()

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Printf(
		"Server started at http://localhost:8080\n",
	)
}
