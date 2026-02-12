package main

import (
	"log"

	"postman/internal/server"
)

func main() {
	srv := server.New()

	log.Println("🚀 Server is running on http://localhost:8080")

	
	if err := srv.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
