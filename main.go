package main

import (
	"fmt"
	"net/http"
	"os"
	"simple_transactional_ledger/config"
	"simple_transactional_ledger/pkg/database"
	"simple_transactional_ledger/routes"
)

func main() {
	config.Getdotenv()
	db :=database.ConnectMySQL()


	server := http.NewServeMux()
	routes.MapRoutes(server, db)

	fmt.Printf("Server is running on port %s\n", os.Getenv("APP_URL")+":"+os.Getenv("APP_PORT"))
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), server)

}