package post

type PostRepo interface {
	InsertPost(post Post) string
}
