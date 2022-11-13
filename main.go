package main

import (
	"awesomeProject1/backend"
	_ "github.com/lib/pq"
)

func main() {
	backend.HandleRequests()
	backend.InitializeDB()
}
