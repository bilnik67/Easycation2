package api

import (
	"fmt"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", HandleFuncStarter).Methods("GET")

	router.Handle("/book/{id}", BookDownloader).Methods("GET")

	http.ListenAndServe(":"+strconv.Itoa(6969), router)
}

func HandleFuncStarter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("a")
	return
}

var BookDownloader = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	GetBooks(w, r)
	return
})

func GetBooks(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/book/{id}", http.StatusSeeOther)
}
