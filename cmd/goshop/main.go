package main

import (
	"log"
	"net/http"
)

func main() {
	port := "8080"

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
