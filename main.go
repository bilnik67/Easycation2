package main

import (
	"awesomeProject1/api"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	api.HandleRequests()
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "db", 5555, "root", "root", "EasyCation_DB")

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
