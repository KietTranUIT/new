package main

import (
	"fmt"
	"net/http"
)

func Show(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Show)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	fmt.Println("Server listening on port 8080")
	server.ListenAndServe()
}
