package Handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, "Welcome to the ecommerce website")
}
