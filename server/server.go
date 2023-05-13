package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	port      string
	mux       *http.ServeMux
	endpoints []string
	handlers  []func(w http.ResponseWriter, r *http.Request)
}

func Start() {
	server := &Server{
		port:      ":3000",
		mux:       http.NewServeMux(),
		endpoints: []string{"/"},
		handlers:  []func(w http.ResponseWriter, r *http.Request){handleRequest},
	}

	server.bindEndpoints()
	http.ListenAndServe(server.port, server.mux)
}

func (s *Server) bindEndpoints() {
	for i, _ := range s.endpoints {
		s.mux.HandleFunc(s.endpoints[i], s.handlers[i])
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The server is live")
	fmt.Printf("Request made to: %s\n", r.URL)
}
