package main

import (
	"fmt"
	"log"
	"os"

	"parser/api/handlers"
	"parser/db/mongodb"
	"parser/router"

	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
      log.Fatal("Error loading .env file")
    }

    _, dbError := mongodb.ConnectMongoDB(os.Getenv("DATABASE_URL"))
    if dbError != nil {
        log.Fatal(dbError)
    }
    fmt.Println("Connected to MongoDB!")
    
    e := router.New()

    go handlers.PoolingNewBlocks()
    
    e.Start(":8000")


    defer func() {
      if err := mongodb.DisconnectMongoDB(); err != nil {
          log.Fatal("Error disconnecting from MongoDB:", err)
      }
      fmt.Println("Disconnected from MongoDB")
  }()

  select {}
}