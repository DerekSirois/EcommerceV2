package internal

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error   bool
	Message string
	Data    any
}

func sendJson(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = w.Write(jsonData)
	return err
}

func readJson(r http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	return err
}

func errorJson(w http.ResponseWriter, statusCode int, err error) {
	payload := jsonResponse{
		Error:   true,
		Message: err.Error(),
	}
	_ = sendJson(w, statusCode, payload)
}
