package main

import (
	"log"
	"net/http"
	"os"
)

func bindAddr() string {
	if value := os.Getenv("PORT"); value != "" {
		return ":" + value
	}
	if value := os.Getenv("BIND_ADDR"); value != "" {
		return value
	}
	return ":9000"
}

func main() {
	log.Printf("listening on %s", bindAddr())
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(bindAddr(), nil)
}
