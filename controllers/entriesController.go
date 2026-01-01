package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"simple_transactional_ledger/models"
)

func IndexEntries(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`SELECT * FROM entries`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var entries []models.Entry
		for rows.Next(){
			var e models.Entry
			err := rows.Scan(&e.Id, &e.AccountId, &e.Amount, &e.CreateAt,&e.UpdateAt)
			if err != nil {
				log.Fatal(err)
			}
			entries = append(entries, e)
		}
		
		json.NewEncoder(w).Encode(entries)
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