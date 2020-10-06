package main

import (
	"fmt"
	"github.com/pwalters04/first-gPRC/database"
)

func main (){
	fmt.Println("Testing DB Connection to AWS RDS Postgres")

	database.DBConnection()
}
