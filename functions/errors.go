package functions

import "net/http"

func CheckErr1(s string, w http.ResponseWriter, r *http.Request) {
	if s == "Bad Request" {
		ReturnCode400(w, r)
	} else if s == "File Not Found" {
		w.WriteHeader(http.StatusNotFound)
		ReturnCode404(w, r)
	} else {
		ReturnCode500(w, r)
	}
}
