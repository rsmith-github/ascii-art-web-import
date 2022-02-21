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
	ascii_art_str := functions.MakeMapSimple(input, w, r)

	tmpl, _ := template.ParseFiles("templates/asciiart.html")
	tmpl.Execute(w, ascii_art_str)
}
