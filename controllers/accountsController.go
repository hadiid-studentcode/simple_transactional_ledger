package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"simple_transactional_ledger/models"
	"strconv"
)


func IndexAccounts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var accounts []models.Account

		stmt, err := db.Prepare(`SELECT id,name,balance,create_at,update_at FROM accounts`)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		rows, err := stmt.Query()
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var a models.Account
			err := rows.Scan(&a.Id, &a.Name, &a.Balance, &a.CreateAt, &a.UpdateAt)
			if err != nil {
				log.Fatal(err)
			}
			accounts = append(accounts, a)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(accounts)
	}
}


func ShowAccount(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err	:= strconv.ParseInt(idStr, 10, 64)
		if err != nil{
			http.Error(w,"ID harus berupa angka", http.StatusBadRequest)
			return
		}
		row := db.QueryRowContext(r.Context(),"SELECT id,name,balance,create_at,update_at FROM accounts WHERE id = ?", id)
		if err := row.Err(); err != nil {
			http.Error(w, "Error Executing Query", http.StatusInternalServerError)
			return
		}
		var a models.Account
		err = row.Scan(&a.Id, &a.Name, &a.Balance, &a.CreateAt, &a.UpdateAt)
		if err != nil {
			http.Error(w,"Error Scanning Row",http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(a)
	}
}

func CreateAccount(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Write([]byte("Halaman Create Account"))
			return
		}

		name := r.FormValue("name")
		balanceStr := r.FormValue("balance")

		balance, err := strconv.ParseFloat(balanceStr, 64)
		if err != nil {
			http.Error(w, "Error parsing balance", http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("INSERT INTO accounts (name, balance) VALUES (?, ?)")
		if err != nil {
			http.Error(w, "Error preparing query", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(name, balance)
		if err != nil {
			http.Error(w, "Error creating account", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Account created successfully"))
	}
}

func UpdateAccount(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.Write([]byte("Halaman Update Account"))
			return
		}

		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "ID harus berupa angka", http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare(`
			UPDATE accounts
			SET name = ?, balance = ?
			WHERE id = ?
		`)
		if err != nil {
			http.Error(w, "Error preparing query", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(r.FormValue("name"), r.FormValue("balance"), id)
		if err != nil {
			http.Error(w, "Error update account", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Account updated successfully"))
	}
}


func DeleteAccount(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodDelete {
			w.Write([]byte("Halaman Delete Account"))
			return
		}

		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "ID harus berupa angka", http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("DELETE FROM accounts WHERE id = ?")
		if err != nil {
			http.Error(w, "Error preparing query", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(id)
		if err != nil {
			http.Error(w, "Error deleting account", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
