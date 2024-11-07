package middlewares

import (
	"log"
	"parser/db/mongodb"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)


func MongoDBMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        client := mongodb.GetClient()
        c.Set(mongodb.MongoClientKey, client)
        return next(c)
    }
}

func SetMongoDBMiddleWare(c *echo.Echo) {
	c.Use(MongoDBMiddleware)
}

func GetMongoClient(c echo.Context) *mongo.Client {
    client, ok := c.Get(mongodb.MongoClientKey).(*mongo.Client)
    if !ok {
        log.Fatal("MongoDB client not found in context")
    }
    return client
}