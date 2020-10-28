package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tonymj76/maze/handlers"
	"github.com/tonymj76/maze/mongodb"
)

func main() {
	if err := godotenv.Load("myenv.env"); err != nil {
		log.Fatalf("Could not connect to .env %v", err)
	}
	host := os.Getenv("HOST")
	if host == "" {
		log.Println("$PORT must be set")
	}
	url := os.Getenv("URL")
	if host == "" {
		log.Println("$PORT must be set")
	}
	session, err := mongodb.NewSeasion(url)
	if err != nil {
		log.Fatalf("connection error %v", err)
	}
	server := handlers.Service{Session: session}
	r := handlers.SetupRouter(server)
	r.Run(":" + host)
}
