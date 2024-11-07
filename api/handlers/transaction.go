package handlers

import (
	"context"
	"net/http"
	"parser/api/middlewares"
	"parser/db/mongodb"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type TransactionResponse struct {
    Address     string        `json:"address" bson:"address"`
    Transactions []Transaction `json:"transactions" bson:"transactions"`
}

func GetTransactionByAddress(c echo.Context) error  {
	address := c.Param("address")

	client := middlewares.GetMongoClient(c)
	collection := client.Database(mongodb.DatabaseName).Collection(mongodb.CollectionName)

	var result TransactionResponse
	err:= collection.FindOne(context.Background(), bson.M{"address": address}).Decode(&result)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Address not subscribed"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"address":     result.Address,
		"transactions": result.Transactions,
	})
}