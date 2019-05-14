package main

import (
	"backend/router"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	router := router.Router()
	log.Fatal(http.ListenAndServe(":3333", router))
}
