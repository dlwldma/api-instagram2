package post

type PostStorage interface {
	UploadImages(images64 []string, filename string) ([]string, error)
}
