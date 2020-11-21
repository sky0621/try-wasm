package main

import (
	"log"
	"net/http"
)

func main() {
	port := "9595"
	log.Printf("listen on http://localhost:%s", port)
	http.Handle("/", http.FileServer(http.Dir("public")))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
