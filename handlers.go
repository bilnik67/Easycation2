package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var BookDownloader = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/book/{id}", http.StatusSeeOther)
	return
})

type userInputRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Region   string `json:"region"`
	Country  string `json:"country"`
}

type userInputLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// username
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

	//username := v.Username
	//password, _ := bcrypt.GenerateFromPassword([]byte(v.Password), 10)
	//mail := v.Email
	//region := v.Region
	//country := v.Country
	//AccountCreationDateTime := time.Now().UTC()
	//Globals.DB.Exec( /*SQL*/ )
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

	//username := v.Username
	//password, _ := bcrypt.GenerateFromPassword([]byte(v.Password), 10)
	//LoginDateTime := time.Now().UTC()
	//Globals.DB.Exec( /*SQL*/ )
})
