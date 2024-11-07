package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BlockNumberResponse struct {
    JSONRPC string `json:"jsonrpc"`
    Result  string `json:"result"`
    ID      int    `json:"id"`
}

type Transaction struct {
    Hash     string `json:"hash"`
    From     string `json:"from"`
    To       string `json:"to"`
    Value    string `json:"value"`
    // Add other transaction fields as needed
}

type Block struct {
    Number       string        `json:"number"`
    Transactions []Transaction `json:"transactions"`
}


func GetCurrentBlockNumber(c echo.Context) error {
    blockNumber, err := getCurrentBlockNumber()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": fmt.Sprintf("Failed to get block number: %v", err),
        })
    }

    return c.JSON(http.StatusOK, map[string]string{
        "blockNumber": blockNumber,
    })
}


func GetTransactionsByBlockNumber(c echo.Context) error {
	blockNumber := c.Param("blockNumber")
    transactions, err := getTransactionsByBlockNumber(blockNumber)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": fmt.Sprintf("Failed to get transaction by block number: %v", err),
        })
    }

    return c.JSON(http.StatusOK, map[string][]Transaction{
        "transactions": transactions,
    })
}

