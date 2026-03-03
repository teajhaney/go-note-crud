package db

import (
	"context"
	"fmt"
	"notes-api/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func Connect (config config.Config)(*mongo.Client, *mongo.Database, error) {

	// Create a context with a timeout of 10 seconds for the connection attempt and prevent app from freezing indefinitely if the database is unreachable
	ctx, cancel := context.WithTimeout(context.Background(),10* time.Second)
	defer cancel()

		clientOpts := options.Client().ApplyURI(config.MongoURI)

		// Connect to MongoDB using the provided URI and options
		client, err := mongo.Connect(ctx, clientOpts)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
		}

		// Ping the database to verify the connection is established successfully
		if err := client.Ping(ctx, nil); err != nil {
			return nil, nil, fmt.Errorf("failed to ping MongoDB: %v", err)
		}

		// Access the specified database using the MongoDB client and return both the client and database instances for further operations
		database := client.Database(config.MongoDBName)

		return client, database, nil

}


func Disconnect(client *mongo.Client) error {
	// Create a context with a timeout of 10 seconds for the disconnection attempt and prevent app from freezing indefinitely if the database is unreachable
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Disconnect from MongoDB using the provided client and context, ensuring that resources are released properly
	if err := client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %v", err)
	}

	return nil
}
