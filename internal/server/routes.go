package server

import (
	"net/http"

	"postman/internal/handler"
)

func (s *Server) RegisterRoutes() {
	// Home page
	homeHandler := handler.NewHomeHandler()
	s.mux.Handle("/", homeHandler)

	// Static files
	fs := http.FileServer(http.Dir("web/static"))
	s.mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Send request
	s.mux.HandleFunc("/send-request", handler.SendRequestHandler)
}
