package post

import "os"

type CreatePostDto struct {
	UserId      string
	Description string
	File        *os.File
}
