package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)
func ConnectMySQL() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
	 os.Getenv("DB_USER"), 
	 os.Getenv("DB_PASSWORD"), 
	 os.Getenv("DB_HOST"), 
	 os.Getenv("DB_PORT"), 
	 os.Getenv("DB_NAME"))
	
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}	
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db

}