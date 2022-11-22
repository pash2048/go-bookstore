package routes

import (
	"github.com/gorilla/mux"
	"github.com/pash2048/go-bookstore/pkg/controllers"
)

var RegiterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("Delete")
}
