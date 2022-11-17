package backend

import (
	"database/sql"
	"time"
)

func RegisterUser(db *sql.DB, register userInputRegister) (err error, result sql.Result) {
	rows, err := db.Query(`SELECT "klant_id" FROM "gebruiker" WHERE email=$1`, register.Email)
	if err != nil {
		return err, result
	}
	if rows != nil {
		return err, result
	}

	rows, err = db.Query(`SELECT "klant_id" FROM "gebruiker" WHERE gebruikersnaam=$1`, register.Username)
	if err != nil {
		return err, result
	}
	if rows != nil {
		return err, result
	}

	result2 := result
	err2 := err
	result2, err2 = db.Exec(`insert into "gebruiker"("gebruikersnaam","wachtwoord","email","created_on", "last_login") values($1, $2, $3, $4, $5)`, register.Username, register.Password, register.Email, time.Now().UTC(), time.Now().UTC())
	if err2 != nil {
		return err2, result2
	}
	//sql code goes here
	return err, result2
}

func LoginUser(db *sql.DB, login userInputLogin) (err error, result sql.Result) {
	//result, err = db.Query('SELECT "username","demoname","demonumber","running" FROM "gebruiker" WHERE username=$1 AND demoname=$2 AND demonumber=$3')
	if err != nil {
		return err, result
	}
	return err, result
}
