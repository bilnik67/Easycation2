package backend

import (
	"database/sql"
	"strconv"
	"time"
)

func RegisterUser(db *sql.DB, register userInputRegister) (err error, result sql.Result) {
	rows, err := db.Query(`SELECT "klant_id" FROM "gebruiker" WHERE email=$1 OR gebruikersnaam=$2`, register.Email, register.Username)
	if err != nil {
		return err, result
	}

	defer rows.Close()
	//TODO add salt
	if rows.Next() {
		return err, result
	}
	hashedPassword, err := hash(register.Password)
	if err != nil {
		return
	}
	result, err = db.Exec(`insert into "gebruiker"("gebruikersnaam","wachtwoord","email","created_on", "last_login") values($1, $2, $3, $4, $5)`, register.Username, strconv.Itoa(hashedPassword), register.Email, time.Now().UTC(), time.Now().UTC())

	return err, result
}

// TODO fix login
func LoginUser(db *sql.DB, login userInputLogin) (err error, result sql.Result) {
	//result, err = db.Query('SELECT "username","demoname","demonumber","running" FROM "gebruiker" WHERE username=$1 AND demoname=$2 AND demonumber=$3')
	if err != nil {
		return err, result
	}
	return err, result
}

//TODO fix login through JS
//TODO fix register through JS
