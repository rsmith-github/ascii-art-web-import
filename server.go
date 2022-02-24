package main

import (
	"asciiweb/functions"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// index page
	fileServer := http.FileServer(http.Dir("./templates"))
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
	ascii_art_str, err := functions.MakeMapSimple(input, w, r)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, _ := template.ParseFiles("templates/asciiart.html")
	tmpl.Execute(w, ascii_art_str)
}

// MethodNotAllowed replies to the request with an HTTP 400 Bad Request error.
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "400 Bad Request", http.StatusBadRequest)
}

// InternalServerError replies to the request with an HTTP 500 Internal Server Error error.
func InternalServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
}

// BadRequest replies to the request with an HTTP 404 Not Found Error error.
func BadRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 Bad Request", http.StatusNotFound)
}
