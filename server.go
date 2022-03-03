package main

import (
	"asciiweb/functions"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// index page
	fileServer := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fileServer)
	http.HandleFunc("/ascii-art", functions.FormHandler)

	fmt.Printf("Starting server at http://localhost:8090\n")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
