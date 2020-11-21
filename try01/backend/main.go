package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("asset")))
	if err := http.ListenAndServe(":9595", nil); err != nil {
		log.Fatal(err)
	}
}
