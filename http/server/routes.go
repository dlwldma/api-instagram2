package server

import (
	"instagram2/domains/post"
	"net/http"
)

func (s *Server) NewAppHandler() *http.ServeMux {
	router := http.NewServeMux()

	postModule := post.NewPostModule()
	router.Handle("/post/", http.StripPrefix("/post", postModule.Handler))

	return router
}
