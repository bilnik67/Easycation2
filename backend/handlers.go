package backend

import (
	"encoding/json"
	"fmt"
	_ "github.com/gorilla/securecookie"
	"io/ioutil"
	"net/http"
	"time"
)

var BookDownloader = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/book/{id}", http.StatusSeeOther)
	return
})

type userInputRegister struct {
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	Email            string    `json:"email"`
	Region           string    `json:"region"`
	Country          string    `json:"country"`
	RegisterDateTime time.Time `json:"datetime,omitempty"`
}

type userInputLogin struct {
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	LoginDateTime time.Time `json:"datetime,omitempty"`
}

var RegisterHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	v := &userInputRegister{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
	}

	err = RegisterUser(DB, *v)
	if err != nil {
		return
	}
})

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	v := &userInputLogin{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	err = LoginUser(DB, *v)
	if err != nil {
		return
	}
})
