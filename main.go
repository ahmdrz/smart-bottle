// HooshNoosh project main.go
package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Checking database connection...")
	if OpenDatabase() != nil {
		log.Fatalln("Connection failed !")
	}
	log.Println("Database is connected to web service")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":2514", router))
}
