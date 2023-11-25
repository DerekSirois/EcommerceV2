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

	return router
}
