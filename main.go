package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", strings.Join(os.Environ(), "\n"))
	})

	log.Println("Starting on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
