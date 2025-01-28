package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

type MongoConfig struct {
	UserName string
	Password string
	DBName   string
}

func LoadMongoConfig() (*MongoConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	cfg := &MongoConfig{
		UserName: os.Getenv("MONGODB_USERNAME"),
		Password: os.Getenv("MONGODB_PASSWORD"),
		DBName:   os.Getenv("MONGODB_DBNAME"),
	}

	var missingVars []string
	if cfg.UserName == "" {
		missingVars = append(missingVars, "MONGODB_USERNAME")
	}
	if cfg.Password == "" {
		missingVars = append(missingVars, "MONGODB_PASSWORD")
	}
	if cfg.DBName == "" {
		missingVars = append(missingVars, "MONGODB_DBNAME")
	}
	if len(missingVars) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %v", missingVars)
	}

	return cfg, nil
}

func InitMongoDB() error {
	cfg, err := LoadMongoConfig()
	if err != nil {
		return err
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://%s:%s@demo.veati.mongodb.net/?retryWrites=true&w=majority",
			cfg.UserName, cfg.Password)).SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("error pinging MongoDB: %v", err)
	}

	Db = client.Database(cfg.DBName)
	fmt.Println("Connected to MongoDB!")
	return nil
}

func CloseMongoDB() {
	if Db != nil {
		if err := Db.Client().Disconnect(context.Background()); err != nil {
			panic(fmt.Errorf("error disconnecting from MongoDB: %v", err))
		}
		fmt.Println("Disconnected from MongoDB")
	}
}
