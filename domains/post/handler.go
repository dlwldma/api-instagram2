package post

import (
	"net/http"
)

type PostHandler struct {
	Handler *http.ServeMux
	service *PostService
}

func NewPostHandler(service *PostService) *PostHandler {
	handler := http.NewServeMux()

	return &PostHandler{
		Handler: handler,
		service: service,
	}
}

func (h *PostHandler) InitializeRoutes() *http.ServeMux {
	h.Handler.HandleFunc("GET /hola/", h.getPosts)
	return h.Handler
}
func (h *PostHandler) getPosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.service.Greetings()))
}
