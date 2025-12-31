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

func UpdateAccount(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut{
				w.Write([]byte("Halaman Update Account"))
			} else {
				idStr := r.PathValue("id")
				id, err := strconv.ParseInt(idStr, 10, 64)
				if err != nil {
					http.Error(w, "ID harus berupa angka", http.StatusBadRequest)
					return
				}

				balanceStr := r.FormValue("balance")
				balance,err := strconv.ParseFloat(balanceStr,64)
				
				if err != nil {
					w.Write([]byte("Error parsing balance"))
					return
				}
				accounts := models.Account{
					Id: id,
					Name: r.FormValue("name"),
					Balance: balance,

				}
				_, err = db.Exec(`
					UPDATE accounts
					SET
						name = ?,
						balance = ?
					WHERE id = ?
				`, accounts.Name, accounts.Balance, accounts.Id)

				if err != nil {
					w.Write([]byte("Error update account"))
					return
				}
				w.Write([]byte("Account updated successfully"))
			}
			

			

	}	
}


func DeleteAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Halaman Delete Account"))
	}	
}