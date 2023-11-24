package main

import (
	"gateway/internal"
	"log"
	"net/http"
)

const webPort = ":80"

func main() {
	log.Println("Serving on port ", webPort)
	log.Fatal(http.ListenAndServe(webPort, internal.Routes()))
}
