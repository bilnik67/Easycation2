package backend

import "database/sql"

func RegisterUser(db *sql.DB, register userInputRegister) (err error, result sql.Result) {
	result, err = db.Exec("Select E-mail, Gebruikersnaam From Gebruikers Where E-mail = " + register.Email + " OR Gebruikersnaam = " + register.Username)
	if err != nil {
		return err, result
	}
	if result != nil {
		return err, result
	}

	result2, err2 := db.Exec("VOEG GEBRUIKER TOE")
	if err2 != nil {
		return err2, result2
	}
	//sql code goes here
	return err, result2
}

func LoginUser(db *sql.DB, login userInputLogin) (err error, result sql.Result) {
	result, err = db.Exec("FROM accounts GET * WHERE Email ==" + register.Email + "AND WHERE PASSWORD ==" + login.Password)
	if err != nil {
		return err, result
	}
	return err, result
}
