package mongo

import (
	"context"
	"github.com/dulguundd/logError-lib/errs"
	"github.com/dulguundd/logError-lib/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"restAPIServer/app/dto"
)

type RepositoryDb struct {
	client *mongo.Client
}

func (d RepositoryDb) Query() *errs.AppError {
	// Define the database and collection
	collection := d.client.Database("productCatalog").Collection("ProductOffering")

	// Define filter to find documents with age greater than 30
	filter := bson.M{"lifecycleStatus": "Active"}

	findOptions := options.Find()
	findOptions.SetLimit(10)

	// Count documents
	//count, err := collection.CountDocuments(context.Background(), filter)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("query count: ", count)

	// Define a slice to store the results
	var results []dto.ProductOffering

	// Find documents that match the filter
	cursor, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	count1 := 0
	// Iterate through the cursor and decode each document into a Person object
	for cursor.Next(context.Background()) {
		var productOffering dto.ProductOffering
		if err := cursor.Decode(&productOffering); err != nil {
			log.Fatal(err)
		}
		count1++
		results = append(results, productOffering)
	}
	log.Println("query count in response: ", count1)

	// Print the results
	//fmt.Println("Results:")
	//for _, result := range results {
	//	fmt.Printf("Id: %s, Name: %d, Status: %s\n", result.Id, result.Name, result.LifecycleStatus)
	//}
	//
	// Close the cursor once finished
	err = cursor.Close(context.Background())
	if err != nil {
		return nil
	}

	return nil
}

func (d RepositoryDb) QueryById() (*dto.ProductOffering, *errs.AppError) {
	logger.Info("Id Query worked")

	// Define the database and collection
	collection := d.client.Database("productCatalog").Collection("ProductOffering")

	// Define filter to find documents with age greater than 30
	filter := bson.M{"id": "7692"}

	findOptions := options.Find()
	findOptions.SetLimit(10)

	// Define a slice to store the results
	var results dto.ProductOffering

	// Find documents that match the filter
	err := collection.FindOne(context.Background(), filter).Decode(&results)
	if err != nil {
		log.Fatal(err)
		return nil, errs.NewUnexpectedError(err.Error())
	}

	// Print the results
	//fmt.Println("Results:")
	//for _, result := range results {
	//	fmt.Printf("Id: %s, Name: %d, Status: %s\n", result.Id, result.Name, result.LifecycleStatus)
	//}

	return &results, nil
}

func (d RepositoryDb) QueryByIdFake() (*dto.ProductOffering, *errs.AppError) {
	logger.Info("Rest called")

	// URL of the REST API endpoint
	url := "http://localhost:8081/home"

	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: Unexpected status code: %d", resp.StatusCode)
	}

	return nil, nil
}

func NewRepositoryDb(dbClient *mongo.Client) RepositoryDb {
	return RepositoryDb{dbClient}
}
