package main

import (
	"fmt"
	"net/http"
	"os"
)

func Show(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Show)
	port := os.Getenv("HTTP_PLATFORM_PORT")

	if port == "" {
		port = "8080"
	}

	server := http.Server{
		Addr:    "127.0.0.1:" + port,
		Handler: mux,
	}

	fmt.Println("Server listening on port 8080")
	server.ListenAndServe()
}
