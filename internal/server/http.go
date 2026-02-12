package server

import "net/http"

type Server struct {
	mux *http.ServeMux
}

func New() *Server {
	mux := http.NewServeMux()

	s := &Server{
		mux: mux,
	}

	s.RegisterRoutes()

	return s
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}
