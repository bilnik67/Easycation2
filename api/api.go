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

	http.ListenAndServe(":"+strconv.Itoa(6969), router)
}

func HandleFuncStarter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("a")
	return
}
