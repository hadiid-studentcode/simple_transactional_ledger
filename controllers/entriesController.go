package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"simple_transactional_ledger/models"
	"strconv"
)

func IndexEntries(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stmt, err := db.PrepareContext(r.Context(), `
		SELECT e.id, e.account_id, e.amount, e.create_at, e.update_at, a.name, a.balance
		FROM entries e
		JOIN accounts a ON e.account_id = a.id
		`)
		if err != nil {
			http.Error(w, "Error preparing query", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		rows, err := stmt.QueryContext(r.Context())
		if err != nil {
			http.Error(w, "Error executing query", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var entries []models.Entry
		for rows.Next(){
			var e models.Entry
			err := rows.Scan(&e.Id, &e.AccountId, &e.Amount, &e.CreateAt, &e.UpdateAt, &e.Name, &e.Balance)
			if err != nil {
				log.Fatal(err)
			}
			entries = append(entries, e)
		}
		
		json.NewEncoder(w).Encode(entries)
	}
}

func ShowEntry(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "ID harus berupa angka", http.StatusBadRequest)
			return
		}
		stmt, err := db.PrepareContext(r.Context(), `
		SELECT e.id, e.account_id, e.amount, e.create_at, e.update_at, a.name, a.balance
		FROM entries e
		JOIN accounts a ON e.account_id = a.id
		WHERE e.id = ?`)
		if err != nil {
			http.Error(w, "Error preparing query", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		row := stmt.QueryRowContext(r.Context(), id)

		if err := row.Err(); err != nil {
			http.Error(w,"Error Executing Query",http.StatusInternalServerError)
			return
		}

		var e models.Entry
		err = row.Scan(&e.Id,&e.AccountId, &e.Amount, &e.CreateAt, &e.UpdateAt, &e.Name, &e.Balance)
		if err != nil {
			http.Error(w,"Error Scanning Row",http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(e)
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