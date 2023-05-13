package server

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRequest)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The server is live")
	})

	port := "8080"
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The server is live")
	fmt.Printf("Request made to: %s\n", r.URL)
}
