package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type FormData struct {
	mr *multipart.Reader
}

func NewFormData(r *http.Request) (*FormData, error) {
	mr, err := r.MultipartReader()
	if err != nil {
		log.Println("Error con reader --", err)
		return nil, errors.New("error when creating new multipart.Reader")
	}
	return &FormData{
		mr: mr,
	}, nil
}

func (fd *FormData) GetFormData() (map[string]any, error) {
	var part *multipart.Part
	var err error
	var data = map[string]any{}
	var files = []*os.File{}
	for {
		if part, err = fd.mr.NextPart(); err != nil {
			if err != io.EOF {
				log.Printf("Hit error while fetching next part: %s", err.Error())
				fmt.Println("Error occured during upload")
				return nil, err
			}
			return data, nil
		}

		filename := part.FileName()
		name := part.FormName()

		if filename == "" {
			value, _ := fd.getValueFromPart(part)
			data[name] = value
		} else {
			file := fd.getTemporalFile(part)
			files = append(files, file)
			data[name] = files
		}
	}
}

func (fd *FormData) getValueFromPart(part *multipart.Part) (int, error) {
	var value int
	var err error
	for {
		if value, err = part.Read(make([]byte, 4096)); err != nil {
			if err != io.EOF {
				log.Printf("Hit error while fetching next part: %s", err.Error())
				fmt.Println("Error occured during upload")
				return 0, err
			}
			return value, nil
		}
	}
}

func (fd *FormData) getTemporalFile(part *multipart.Part) *os.File {
	var n int
	var err error

	chunk := make([]byte, 4096)
	var tempfile *os.File
	var filesize int
	var uploaded bool

	tempfile, err = os.CreateTemp(os.TempDir(), "image*.tmp")
	if err != nil {
		fmt.Printf("Hit error while creating temp file: %s", err.Error())
		fmt.Println("Error occured during upload")
		return nil
	}
	log.Printf("Temporary filename: %s\n", tempfile.Name())

	for !uploaded {
		if n, err = part.Read(chunk); err != nil {
			if err != io.EOF {
				log.Printf("Hit error while reading chunk: %s", err.Error())
				fmt.Println("Error occured during upload")
				return nil
			}
			uploaded = true
		}

		if n, err = tempfile.Write(chunk[:n]); err != nil {
			log.Printf("Hit error while writing chunk: %s", err.Error())
			fmt.Println("Error occured during upload")
			return nil
		}
		filesize += n
	}
	_, err = tempfile.Seek(0, 0)
	if err != nil {
		log.Printf("Error resetting file pointer: %s", err.Error())
		fmt.Println("Error occured during upload")
		return nil
	}

	length, _ := tempfile.Stat()
	log.Printf("Uploaded filesize: %d bytes", length.Size())
	defer tempfile.Close()
	defer os.Remove(tempfile.Name())

	return tempfile
}
