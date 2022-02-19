package main

import (
	"asciiweb/functions"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// file Server to host static files
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/ascii-art", formHandler)

	fmt.Printf("Starting server at http://localhost:8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// handler to accept the /ascii-art request
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	input := r.FormValue("input")
	functions.MakeMapSimple(input, w, r)
}
