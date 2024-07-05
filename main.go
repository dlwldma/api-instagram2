package main

import (
	"fmt"
	"instagram2/http/server"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	fmt.Println(os.Getenv("APP_PORT"))
	PORT := os.Getenv("APP_PORT")
	server := server.NewServer(PORT)

	server.Start()
}
