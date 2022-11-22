package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pash2048/go-bookstore/pkg/routes"
	_ "gorm.io/driver/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegiterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9090", r))
}
