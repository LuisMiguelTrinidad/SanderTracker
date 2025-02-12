package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/LuisMiguelTrinidad/Sandertracker/logging"
)

var Db *mongo.Database

type MongoConfig struct {
	UserName string
	Password string
	DBName   string
}

func LoadMongoConfig() (*MongoConfig, error) {
	if err := godotenv.Load(); err != nil {
		logging.SystemFatalLog(fmt.Sprintf("Failed to load .env file: %v", err))
	}

	cfg := &MongoConfig{
		UserName: os.Getenv("MONGODB_USERNAME"),
		Password: os.Getenv("MONGODB_PASSWORD"),
		DBName:   os.Getenv("MONGODB_DBNAME"),
	}

	if cfg.UserName == "" {
		logging.SystemFatalLog(fmt.Sprintf("missing variable MONGODB_USERNAME from environment"))
	}
	if cfg.Password == "" {
		logging.SystemFatalLog(fmt.Sprintf("missing variable MONGODB_PASSWORD from environment"))
	}
	if cfg.DBName == "" {
		logging.SystemFatalLog(fmt.Sprintf("missing variable MONGODB_DBNAME from environment"))
	}
	logging.SystemInfoLog("Loaded MongoDB config from environment variables")
	return cfg, nil
}

func InitMongoDB() {
	cfg, err := LoadMongoConfig()
	if err != nil {
		logging.SystemFatalLog(fmt.Sprintf("Failed to load MongoDB config: %v", err))
		panic("")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://%s:%s@demo.veati.mongodb.net/?retryWrites=true&w=majority", cfg.UserName, cfg.Password),
	).SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		logging.SystemFatalLog(fmt.Sprintf("Failed to connect to MongoDB: %v", err))
	}

	if err = client.Ping(ctx, nil); err != nil {
		logging.SystemFatalLog(fmt.Sprintf("Failed to ping MongoDB: %v", err))
	}

	Db = client.Database(cfg.DBName)
	logging.SystemInfoLog("Connected to MongoDB")
}

func CloseMongoDB() {
	if Db != nil {
		if err := Db.Client().Disconnect(context.Background()); err != nil {
			logging.SystemFatalLog(fmt.Sprintf("error disconnecting from MongoDB: %v", err))
		}
		logging.SystemInfoLog("Disconnected from MongoDB")
	}
}
