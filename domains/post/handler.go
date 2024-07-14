package post

import (
	"encoding/json"
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
	var formData *utils.FormData
	var formValues map[string]any
	var response *utils.HttpResponse
	var err error

	if formData, err = utils.NewFormData(r); err != nil {
		utils.HandleError(w, err)
		return
	}

	if formValues, err = formData.GetFormData(); err != nil {
		utils.HandleError(w, err)
		return
	}

	jsonbody, _ := json.Marshal(formValues)
	dto := CreatePostDto{}
	if err = json.Unmarshal(jsonbody, &dto); err != nil {
		utils.HandleError(w, err)
		return
	}
	if _, err = utils.AreAllFieldsSet(dto); err != nil {
		utils.HandleError(w, err)
		return
	}
	log.Println("POST: create-post body...", dto)

	//createPost := h.service.CreatePost(dto)
	response = utils.CreateNewResponse(200, "success", nil)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Println("POST: create-post ended()...")
}
