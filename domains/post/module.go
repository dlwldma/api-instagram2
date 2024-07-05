package post

import (
	"instagram2/db"
	"net/http"
)

type PostModule struct {
	Handler *http.ServeMux
}

func NewPostModule() *PostModule {
	client := db.NewMongoDb()
	repo := NewMongoPostRepo(client)
	service := NewPostService(repo)
	handler := NewPostHandler(service)

	return &PostModule{
		Handler: handler.InitializeRoutes(),
	}
}
