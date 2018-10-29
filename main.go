package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/durban89/wiki/router"
)

func main() {
	// 路由
	router.Routes()

	err := http.ListenAndServe(":8090", nil)

	if err != nil {
		log.Println("Server Start Failed")
	} else {
		fmt.Println("Listening on 0.0.0.0:8090")
		log.Println("Listening on 0.0.0.0:8090")
	}
}
