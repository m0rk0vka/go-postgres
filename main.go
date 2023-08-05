package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/m0rk0vka/go-postgres/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server in the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
