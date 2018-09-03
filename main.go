package main

import (
	"log"
	"net/http"

	"github.com/durban89/wiki/router"
)

func main() {
	router.Routes()

	err := http.ListenAndServe(":8090", nil)

	if err != nil {
		log.Println("Server Start Failed")
	} else {
		log.Println("Listening on 0.0.0.0:8090")
	}
}
