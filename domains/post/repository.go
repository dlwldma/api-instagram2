package post

type PostRepo interface {
	GetPosts() []string
}
