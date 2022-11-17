package main

import (
	"awesomeProject1/backend"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	CheckError(backend.InitializeDB())
	backend.HandleRequests()
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
