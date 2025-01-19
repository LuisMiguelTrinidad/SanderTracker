package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongoDB() (*mongo.Client, error) {
	// Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// Get environment variables
	password := os.Getenv("MONGODB_PASSWORD")
	dbName := os.Getenv("MONGODB_DBNAME")
	if password == "" {
		return nil, fmt.Errorf("MONGODB_PASSWORD environment variable not set")
	}

	// Configure connection options
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://root:%s@demo.veati.mongodb.net/?retryWrites=true&w=majority&appName=%s",
			password, dbName)).SetServerAPIOptions(serverAPI)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	// Test the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(fmt.Errorf("error pinging MongoDB: %v", err))
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

func closeMongoDB(client *mongo.Client) {
	if client != nil {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(fmt.Errorf("error disconnecting from MongoDB: %v", err))
		}
		fmt.Println("Disconnected from MongoDB")
	}
}

func main() {
	client, err := initMongoDB()
	if err != nil {
		panic(err)
	}
	defer closeMongoDB(client)

	// Your application logic here
}
