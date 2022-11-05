package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mayukorin/ebook-merge/app/interfaces/api"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(fmt.Errorf("failed open .env. %w", err))
	}

	firebaseServiceAccountKeyPath := os.Getenv("FIREBASE_SERVICE_ACCOUNT_KEY_PATH")
	fmt.Println(firebaseServiceAccountKeyPath)

	s, err := api.NewServer(os.Getenv("PLANET_SCALE_DSN"))
	if err != nil {
		fmt.Println(fmt.Errorf("failed to new server: %s", err))
	}

	s.Run(os.Getenv("PORT"))

}
