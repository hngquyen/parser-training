package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MongoClientKey = "MongoClient"
const DatabaseName = "parser"

var client *mongo.Client

func ConnectMongoDB(uri string) (*mongo.Client, error) {
    clientOptions := options.Client().ApplyURI(uri)
    var err error
    client, err = mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        return nil, err
    }

    err = client.Ping(context.Background(), nil)
    if err != nil {
        return nil, err
    }
    return client, nil
}

func GetClient() *mongo.Client {
    return client
}

func DisconnectMongoDB() error {
    return client.Disconnect(context.Background())
}