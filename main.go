package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type envs struct {
	items env_item
}
type env_item struct {
	key   string `json:"key"`
	value string `json:"value"`
}

func main() {
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", strings.Join(os.Environ(), "\n"))
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
