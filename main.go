package main

import (
	"log"

	"parser/router"

	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
      log.Fatal("Error loading .env file")
    }
    
    e := router.New()

    e.Start(":8000")
}