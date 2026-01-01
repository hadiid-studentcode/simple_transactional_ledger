package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

func CreateEntry(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Write([]byte("Halaman Create Entrie"))
			return
		}

		accountId, err := strconv.ParseInt(r.FormValue("account_id"), 10, 64)
		if err != nil {
			http.Error(w, "ID harus berupa angka", http.StatusBadRequest)
			return
		}

		amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
		if err != nil {
			http.Error(w, "Error parsing amount", http.StatusBadRequest)
			return
		}

		stmt, err := db.PrepareContext(r.Context(), `
			INSERT INTO entries (account_id, amount)
			VALUES (?, ?)
			
		`)

		if err != nil {
			http.Error(w,"Error preparing query", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		result, err := stmt.ExecContext(r.Context(), accountId, amount)
		if err != nil {
			http.Error(w,"Error Executing Query",http.StatusInternalServerError)
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w,"Error getting last insert id",http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("Entry created successfully with ID %d", id)))
	}
}
func UpdateEntry(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.Write([]byte("Halaman Update Entrie"))
			return
		}

		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "ID harus berupa angka", http.StatusBadRequest)
			return
		}

		accountIdStr := r.FormValue("account_id")
		accountId, err := strconv.ParseInt(accountIdStr, 10, 64)
		if err != nil {
			http.Error(w, "ID harus berupa angka", http.StatusBadRequest)
			return
		}

		amountStr := r.FormValue("amount")
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			
			http.Error(w, "Error parsing amount", http.StatusBadRequest)
			return
		}

		stmt, err := db.PrepareContext(r.Context(), `
			UPDATE entries
			SET account_id = ?, amount = ?
			WHERE id = ?
		`)
		if err != nil {
			http.Error(w,"Error preparing query", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		result, err := stmt.ExecContext(r.Context(), accountId, amount, id)
		if err != nil {
			http.Error(w,"Error Executing Query",http.StatusInternalServerError)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			http.Error(w,"Error getting rows affected",http.StatusInternalServerError)
			return
		}

		if rowsAffected == 0 {
			http.Error(w,"Entry not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}	
}

func DeleteEntry(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.Write([]byte("Halaman Delete Entrie"))
			return
		}

		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "ID harus berupa angka", http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("DELETE FROM entries WHERE id = ?")
		if err != nil {
			http.Error(w, "Error preparing query", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		result, err := stmt.ExecContext(r.Context(), id)
		if err != nil {
			http.Error(w, "Error deleting entry", http.StatusInternalServerError)
			return
		}

		_, err = result.RowsAffected()
		if err != nil {
			http.Error(w, "Error getting rows affected", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
