package controllers

import "net/http"

func IndexEntries() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Halaman Entries"))
	}
}

func ShowEntry() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Halaman Detail Entry"))
	}	
}

func CreateEntry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Halaman Create Entry"))
	}	
}

func UpdateEntry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Halaman Update Entry"))
	}	
}

func DeleteEntry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Halaman Delete Entry"))
	}	
}