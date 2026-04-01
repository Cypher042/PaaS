package database

import (
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect(MONGODB_URI string) *mongo.Database {

	clientOptions := options.Client().ApplyURI(MONGODB_URI)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatalf("error connecting to MongoDB: %v", err)
	}

	// defer client.Disconnect(context.TODO())

	return client.Database("PaaS")

}
