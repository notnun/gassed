package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Printf("[ERROR] %s", err)
		os.Exit(1)
	}
}

func run() error {
	return http.ListenAndServe(":8080", http.HandlerFunc(handleIndex))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You've got gas!")
}
