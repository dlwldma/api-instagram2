package post

import (
	"os"
)

type PostStorage interface {
	UploadImage(file *os.File, filename string) ([]string, error)
}
