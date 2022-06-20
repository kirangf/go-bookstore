package routes

import (
	"github.com/gorilla/mux"
	"github.com/kirangf/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.GetBookById).Methods("GET")
}
