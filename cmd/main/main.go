package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aditya-9901/bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server statred at :9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}
