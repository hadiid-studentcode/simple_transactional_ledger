package routes

import (
	"database/sql"
	"net/http"
	"simple_transactional_ledger/controllers"
)

func MapRoutes(server *http.ServeMux , db *sql.DB){
	server.Handle("/",http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello, World!"))
	}))

	server.HandleFunc("/home", controllers.IndexHome())
	server.HandleFunc("/accounts", controllers.IndexAccounts(db))
	server.HandleFunc("/accounts/create", controllers.CreateAccount(db))
	server.HandleFunc("/accounts/update/{id}", controllers.UpdateAccount(db))
	server.HandleFunc("/accounts/{id}", controllers.ShowAccount(db))
	server.HandleFunc("/accounts/delete/{id}", controllers.DeleteAccount(db))
	server.HandleFunc("/entries", controllers.IndexEntries())
}