package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	fmt.Printf("Welcome to the Nomad LNL App!\n")
	fmt.Printf("Version: " + version + "\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
