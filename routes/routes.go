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

	server.HandleFunc("/entries", controllers.IndexEntries(db))
	server.HandleFunc("/entries/{id}", controllers.ShowEntry(db))
	server.HandleFunc("/entries/create", controllers.CreateEntry(db))
	server.HandleFunc("/entries/update/{id}", controllers.UpdateEntry(db))
	server.HandleFunc("/entries/delete/{id}", controllers.DeleteEntry(db))
}