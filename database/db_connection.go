package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "gprc-db.cbrumqsqxsul.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "*********"
	dbname   = "********"
)

func DBConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Pinging")
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Pinging-Done")
	fmt.Println("Successfully connected!")

	return db
}
