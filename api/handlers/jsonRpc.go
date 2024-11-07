package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"parser/db/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func getSubscribedAddress() ([]string, error) {
	collection := mongodb.GetClient().Database(mongodb.DatabaseName).Collection(mongodb.CollectionName)

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var subscribedAddresses []string
	for cursor.Next(context.Background()) {
		var address struct {
			Address string `bson:"address"`
		}
		if err := cursor.Decode(&address); err != nil {
			return nil, err
		}
		subscribedAddresses = append(subscribedAddresses, address.Address)
	}

	return subscribedAddresses, nil
}

func processingTransactionsForSubscribedAddress(transactions []Transaction) error {
	collection := mongodb.GetClient().Database(mongodb.DatabaseName).Collection(mongodb.CollectionName)

	subscribedAddresses, err := getSubscribedAddress()
	if err != nil {
		return  err
	}

	for _, tx := range transactions {
		for _, address := range subscribedAddresses {
			if tx.From == address || tx.To == address {
				
				_, err := collection.UpdateOne(
					context.Background(),
					bson.M{"address": address},
					bson.M{
						"$push": bson.M{"transactions": tx},
					},
				)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func PoolingNewBlocks() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	var lastBlockNumber string
	for {
		select {
		case <-ticker.C:
			blockNumber, err := getCurrentBlockNumber()
			if err != nil {
				log.Println("Error fetching block number:", err)
				continue
			}

			if blockNumber != lastBlockNumber {
				lastBlockNumber = blockNumber

				transactions, err := getTransactionsByBlockNumber(blockNumber)

				if err != nil {
					log.Println("Error fetching transactions:", err)
					continue
				}

				err = processingTransactionsForSubscribedAddress(transactions)
				if err != nil {
					log.Println("Error processing transactions:", err)
				}
			}
		}
	}
}