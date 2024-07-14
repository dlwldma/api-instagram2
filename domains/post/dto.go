package post

import "os"

type CreatePostDto struct {
	UserId      string
	Description string
	Files       []*os.File
}
