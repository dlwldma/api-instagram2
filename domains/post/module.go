package post

import (
	"instagram2/api/supabase"
	"instagram2/db"
	"net/http"
)

type PostModule struct {
	Handler *http.ServeMux
}

func NewPostModule() *PostModule {
	storage := supabase.NewSupabaseClient()
	collection := db.NewMongoDb("post")
	repo := NewMongoPostRepo(collection)
	service := NewPostService(repo, storage)
	handler := NewPostHandler(service)

	return &PostModule{
		Handler: handler.InitializeRoutes(),
	}
}
