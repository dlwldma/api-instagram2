package post

import (
	"encoding/json"
	"fmt"
	"log"
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
	log.Println("POST: create-post started()...")

	mr, _ := r.MultipartReader()
	formData := utils.NewFormData(mr)

	dto, err := formData.GetFormData()
	if err != nil {
		fmt.Println("No se pudieron obtener los datos del form data")
	}

	log.Println("POST: create-post body...", dto)

	//createPost := h.service.CreatePost(dto)
	response := utils.CreateNewResponse(200, "success", nil)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
