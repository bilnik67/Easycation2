package backend

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func InitializeDB() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "db", 5555, "root", "root", "EasyCation_DB")

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	DB = db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
