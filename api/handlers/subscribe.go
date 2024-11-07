package handlers

import (
	"context"
	"net/http"
	"parser/api/middlewares"
	"parser/db/mongodb"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SubscribeAddress(c echo.Context) error {
	address := c.FormValue("address")
	if address == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Address form field is required"})
    }

	client := middlewares.GetMongoClient(c)
	collection := client.Database(mongodb.DatabaseName).Collection(mongodb.CollectionName)
	
	err := collection.FindOne(context.Background(), bson.M{"address": address,}).Err()
	if err != mongo.ErrNoDocuments {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Address already subscribed"})
	}

	newAddress := bson.M{
		"address":     address,
		"transactions": []Transaction{}, 
	}

	_, err = collection.InsertOne(context.Background(), newAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to subscribe address"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Address subscribed successfully",
		"address": address,
	})
}