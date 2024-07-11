package post

import (
	"encoding/json"
	"fmt"
	"net/http"

	"instagram2/http/utils"
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
	h.Handler.HandleFunc("POST /create-post/", h.createPost)
	return h.Handler
}

func (h *PostHandler) createPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST: create-post started()...")

	decoder := json.NewDecoder(r.Body)
	var body CreatePostDto
	err := decoder.Decode(&body)
	if err != nil {
		fmt.Println("Error con las propiedades")
	}

	createPost := h.service.CreatePost(body)

	response := utils.CreateNewResponse(200, "success", createPost)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
