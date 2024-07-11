package supabase

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	storage_go "github.com/supabase-community/storage-go"
)

type Supabase struct {
	Storage *storage_go.Client
}

func NewSupabaseClient() *Supabase {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	STORAGE_URL := os.Getenv("SUPABASE_STORAGE_URL")
	API_SECRET_KEY := os.Getenv("SUPABASE_PRIVATE_KEY")

	headers := map[string]string{
		"Content-Type": "image/png",
	}
	storageClient := storage_go.NewClient(STORAGE_URL, API_SECRET_KEY, headers)
	return &Supabase{
		Storage: storageClient,
	}
}
