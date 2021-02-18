package model

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "golang_app_data"
)

var dbConnection *sql.DB

// Connect  connect to postgres database
func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	connection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Println("DB connection established..")
	dbConnection = connection
	return connection

}
