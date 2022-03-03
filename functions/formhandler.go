package functions

import (
	"fmt"
	"net/http"
	"text/template"
)

// handler to accept the /ascii-art request
func FormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if r.Method == http.MethodPost {
		input := r.FormValue("input")
		ascii_art_str, err := MakeMapSimple(input, w, r)
		if err != nil {
			CheckErr1(err.Error(), w, r)
			return
		}

		tmpl, err := template.ParseFiles("templates/asciiart.html")
		if err != nil {
			fmt.Println("Error Occurred:")
			fmt.Println(err)
			fmt.Println()
			ReturnCode500(w, r)
			return
		}
		tmpl.Execute(w, ascii_art_str)
	} else {
		tmpl, err := template.ParseFiles("templates/asciiart.html")
		if err != nil {
			fmt.Println("Error Occurred:")
			fmt.Println(err)
			fmt.Println()
			ReturnCode500(w, r)
			return
		}
		tmpl.Execute(w, "")
	}
}
