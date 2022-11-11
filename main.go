package main

import (
	"awesomeProject1/Globals"
	_ "github.com/lib/pq"
)

func main() {
	HandleRequests()
	Globals.InitializeDB()
}
