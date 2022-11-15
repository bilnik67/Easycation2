package backend

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/securecookie"
	"net/http"
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

var s = securecookie.New(securecookie.GenerateRandomKey(50), securecookie.GenerateRandomKey(32))

func SetAccountCookie(w http.ResponseWriter, inputValue string) {
	//make securecookie instance, first value how to encrypt name, second to encrypt value

	value := map[string]string{
		"AccountId": inputValue,
	}
	if encoded, err := s.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:     "cookie",
			Value:    encoded,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
}

func ReadAccountCookie(w http.ResponseWriter, r *http.Request) string {
	if cookie, err := r.Cookie("cookie"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("cookie", cookie.Value, &value); err == nil {
			return value["AccountId"]
		}
	}
	return ""
}
