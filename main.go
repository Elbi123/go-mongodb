package main

import (
	"fmt"
	"log"

	"github.com/Elbi123/gomongo/server"
)

func main() {
	// database.DataStore()
	log.Fatal(server.ServeAPI())
	fmt.Println("go mongodb")
}
