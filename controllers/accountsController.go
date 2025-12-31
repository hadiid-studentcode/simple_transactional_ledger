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
		rows, err := db.Query(`SELECT id,name,balance,create_at,update_at FROM accounts`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var accounts []models.Account
		for rows.Next(){
			var a models.Account

			err := rows.Scan(&a.Id, &a.Name, &a.Balance, &a.CreateAt, &a.UpdateAt)

			if err != nil {
				log.Fatal(err)
			}
			accounts = append(accounts, a)
		}

		if err := rows.Err(); err != nil{
			log.Fatal(err)
		}

		 json.NewEncoder(w).Encode(accounts)
	}
}

func ShowAccount() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Halaman Detail Account"))
	}	
}

func CreateAccount(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	 if r.Method != http.MethodPost{
		w.Write([]byte("Halaman Create Account"))
	 } else {
		balanceStr := r.FormValue("balance")
	 	balance, err := strconv.ParseFloat(balanceStr, 64)
	 	if err != nil {
		 	w.Write([]byte(" Error parsing balance"))
		 	return
		 }

	 	accounts := models.Account{
		Name: r.FormValue("name"),
		Balance: balance,
	 	}

		

		 _, err = db.Exec("INSERT INTO accounts (name, balance) VALUES (?, ?)", accounts.Name, accounts.Balance)
	 	if err != nil {
			 w.Write([]byte("Error creating account"))
		 	return
	 	}
	 
	 	w.Write([]byte("Account created successfully"))
	}
	}	
}

func UpdateAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Halaman Update Account"))
	}	
}

func DeleteAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Halaman Delete Account"))
	}	
}