package post

type PostService struct {
	repo PostRepo
}

func NewPostService(repo PostRepo) *PostService {
	return &PostService{
		repo,
	}
}

func (s *PostService) Greetings() string {
	return "Hola humano desde service!!! :) "
}
