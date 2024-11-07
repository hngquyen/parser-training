package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type JSONRPCRequest struct {
    JSONRPC string      `json:"jsonrpc"`
    Method  string      `json:"method"`
    Params  interface{} `json:"params"`
    ID      int         `json:"id"`
}

type JSONRPCResponse struct {
    JSONRPC string          `json:"jsonrpc"`
    Result  json.RawMessage `json:"result"`
    ID      int             `json:"id"`
}

func sendJSONRPCRequest(method string, params interface{}) (json.RawMessage, error) {
    requestBody, err := json.Marshal(JSONRPCRequest{
        JSONRPC: "2.0",
        Method:  method,
        Params:  params,
        ID:      1,
    })
    if err != nil {
        return nil, err
    }

    resp, err := http.Post(os.Getenv("ALCHEMY_NODE_SERVICES"), "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result JSONRPCResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    return result.Result, nil
}

func getCurrentBlockNumber() (string, error) {
    rawResult, err := sendJSONRPCRequest("eth_blockNumber", []interface{}{})
    if err != nil {
        return "", err
    }

    var blockNumber string
    if err := json.Unmarshal(rawResult, &blockNumber); err != nil {
        return "", err
    }
    return blockNumber, nil
}

func getTransactionsByBlockNumber(blockNumber string) ([]Transaction, error) {
    params := []interface{}{blockNumber, true} 
    rawResult, err := sendJSONRPCRequest("eth_getBlockByNumber", params)
    if err != nil {
        return nil, err
    }


    var block Block
    if err := json.Unmarshal(rawResult, &block); err != nil {
        return nil, err
    }
    return block.Transactions, nil
}