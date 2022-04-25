package main

import (
	"fmt"
	"go-crud/database"
	"go-crud/router"
	"log"
	"net/http"
)

func main() {
	r := router.SetupRouter()

	database.Connect()

	fmt.Println("Starting server on the port 8889...")

	log.Fatal(http.ListenAndServe(":8889", r))
}
