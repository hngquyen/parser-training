package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SubscribeAddress(c echo.Context) error {
	address := c.FormValue("address")
	log.Printf("Address: %#v", address)

	if address == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Address form field is required"})
    }

	// TODO: Add address to database

	return c.JSON(http.StatusOK, map[string]string{
		"address": address,
	})
}