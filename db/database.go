package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() {
	databaseUser := os.Getenv("DB_USER")
	databasePassword := os.Getenv("DB_PASSWORD")
	//databaseUri := os.Getenv("DB_URI")
	databasePort := os.Getenv("DB_PORT")
	//mongoApplyURI := fmt.Sprintf("mongodb://%v:%v@%v:%v/?authSource=admin&authMechanism=SCRAM-SHA-256", databaseUser, databasePassword, databaseUri, databasePort)
	// Change next line to upper commented lines if not use docker containers:
	mongoApplyURI := fmt.Sprintf("mongodb://%v:%v@mongo:%v", databaseUser, databasePassword, databasePort)
	clientOptions := options.Client().ApplyURI(mongoApplyURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Connection error with MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Cannot verify connection with MongoDB:", err)
	}

	DB = client
	log.Println("MongoDB connection successfully")
}

func GetCollection(collectionName string) *mongo.Collection {
	databaseName := os.Getenv("DB_NAME")
	return DB.Database(databaseName).Collection(collectionName)
}
