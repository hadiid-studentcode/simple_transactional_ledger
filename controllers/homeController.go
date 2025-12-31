package controllers

import "net/http"


func IndexHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Halaman Home"))
	}

}