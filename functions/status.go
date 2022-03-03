package functions

import (
	"net/http"
)

func ReturnCode500(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	return
}

func ReturnCode400(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "400 Bad Request", http.StatusBadRequest)
	return
}

func ReturnCode404(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 Not Found", http.StatusNotFound)
	return
}
