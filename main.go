package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/durban89/wiki/router"
)

func main() {
	// 路由
	router.Routes()

	var addr = flag.String("addr", ":8090", "The addr of the application")

	log.Println("Starting web server on", *addr)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
