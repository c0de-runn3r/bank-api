package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := processENV(); err != nil {
		log.Fatal(err)
	}
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}

func processENV() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file")
	}
	return nil
}
