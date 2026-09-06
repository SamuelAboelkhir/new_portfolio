package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type apiConfig struct {
	port         string
	filePathRoot string
	publicPath   string
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	cfg := apiConfig{
		port:         os.Getenv("PORT"),
		filePathRoot: os.Getenv("FILE_PATH"),
		publicPath:   os.Getenv("PUBLIC_PATH"),
	}

	server(&cfg)
}
