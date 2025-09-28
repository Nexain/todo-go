package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database
var TodoCollection *mongo.Collection

func InitMongo() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Mongo connect error:", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Mongo ping error:", err)
	}

	Client = client
	DB = client.Database(dbName)
	log.Println("✅ MongoDB connected:", dbName)
}

func InitCollection() {
	TodoCollection = DB.Collection("todos")
	log.Println("✅ Collection initialized: todos")
}
