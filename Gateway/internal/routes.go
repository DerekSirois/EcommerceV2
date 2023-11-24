package internal

import (
	"gateway/internal/Handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func Routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", Handlers.Index).Methods("GET")

	return router
}
