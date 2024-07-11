package post

import (
	"time"

	"github.com/google/uuid"
)

type PostService struct {
	repo    PostRepo
	storage PostStorage
}

func NewPostService(repo PostRepo, storage PostStorage) *PostService {
	return &PostService{
		repo, storage,
	}
}

func (s *PostService) CreatePost(body CreatePostDto) string {
	time_now := time.Now()
	post_id := uuid.New().String()
	filename := s.CreateFilename(post_id, body.UserId, &time_now)
	img_url := s.EncodeImage64(body.ImageUris64, filename)
	post := &Post{
		Id:          post_id,
		UserId:      body.UserId,
		Date:        time_now,
		Description: body.Description,
		Image_urls:  img_url,
		Likes:       nil,
		Comments:    nil,
	}
	s.repo.InsertPost(*post)
	return "Post creado con exito!"
}

func (s *PostService) EncodeImage64(image_uris []string, filename string) []string {
	image_urls, _ := s.storage.UploadImages(image_uris, filename)
	return image_urls
}

func (s *PostService) CreateFilename(post_id string, user_id string, time *time.Time) string {
	img_path := user_id + "/" + post_id + "/"
	return img_path
}
