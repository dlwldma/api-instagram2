package post

import (
	"instagram2/db"
	"net/http"
)

type PostModule struct {
	Handler *http.ServeMux
}

func NewPostModule() *PostModule {
	collection := db.NewMongoDb("post")
	repo := NewMongoPostRepo(collection)
	service := NewPostService(repo)
	handler := NewPostHandler(service)

	return &PostModule{
		Handler: handler.InitializeRoutes(),
	}
}
