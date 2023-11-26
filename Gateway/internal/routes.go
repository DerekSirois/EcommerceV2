package internal

import (
	"gateway/internal/Handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func Routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", Handlers.Index).Methods("GET")
	router.HandleFunc("/register", Handlers.Register).Methods("POST")
	router.HandleFunc("/login", Handlers.Login).Methods("POST")

	router.HandleFunc("/product", Handlers.GetAllProduct).Methods("GET")
	router.HandleFunc("/product/available", Handlers.GetAllAvailableProduct).Methods("GET")
	router.HandleFunc("/product/outofstock", Handlers.GetAllOutOfStockProduct).Methods("GET")
	router.HandleFunc("/product", Handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/product/{id:[0-9]+}", Handlers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{id:[0-9]+}", Handlers.DeleteProduct).Methods("DELETE")

	return router
}
