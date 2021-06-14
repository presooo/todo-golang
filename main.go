package main

import (
	"fmt"
	"log"
	"net/http"
	"todo/src/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting the server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
