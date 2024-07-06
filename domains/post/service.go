package post

type PostService struct {
	repo PostRepo
}

func NewPostService(repo PostRepo) *PostService {
	return &PostService{
		repo,
	}
}

func (s *PostService) CreatePost(post Post) string {
	s.repo.InsertPost(post)
	return "Post creado con exito!"
}
