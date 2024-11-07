package api

import (
	"parser/api/handlers"

	"github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
    e.GET("/blockNumber", handlers.GetCurrentBlockNumber)
    e.GET("/transactions/:blockNumber", handlers.GetTransactionsByBlockNumber)

	e.POST("/subscribe", handlers.SubscribeAddress)

	e.GET("/transactions/:address", handlers.GetTransactionByAddress)
}