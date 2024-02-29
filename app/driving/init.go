package driving

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func GetDbClientMongo() *mongo.Client {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://databaseAdmin:databaseAdmin123456@10.13.20.3:30017").SetMaxPoolSize(1000).SetMaxConnecting(10)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
