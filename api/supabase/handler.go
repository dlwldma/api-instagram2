package supabase

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

func (s *Supabase) UploadImages(images64 []string, filename string) ([]string, error) {
	var urls []string
	for i := 0; i < len(images64); i++ {
		image64 := images64[i]
		fileBodyUnbased, _ := base64.StdEncoding.DecodeString(image64)
		fileBody := bytes.NewReader(fileBodyUnbased)
		filepath, err := s.Storage.UploadFile("post", filename+strconv.Itoa(i)+".png", fileBody)
		if err != nil {
			fmt.Println("Err ---", err)
			return nil, err
		}

		fmt.Println("Filepath ---", filepath)
		STORAGE_URL := os.Getenv("SUPABASE_STORAGE_URL")
		url := STORAGE_URL + "/object/public/" + filepath.Key
		urls = append(urls, url)
	}
	fmt.Println("Url list ---", urls)
	return urls, nil
}
