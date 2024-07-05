package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Port   string
	Server *http.Server
}

func NewServer(port string) *Server {
	return &Server{
		Port: port,
		Server: &http.Server{
			Addr: ":" + port,
		},
	}
}

func (s *Server) Start() {
	s.Server.Handler = s.NewAppHandler()
	fmt.Println("Server started on port: " + s.Port)
	s.Server.ListenAndServe()
}
