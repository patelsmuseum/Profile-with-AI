package database

import (
	"context"
	"errors"
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDb *mongo.Client
var err error

func InitDB() {
	url := os.Getenv("MONGO_URL")

	mongoDb, err = ConnectMongoDB(url)
	if err != nil {
		logrus.Fatal(err)
	}
	log.Info("MongoDb is connected to "+url+" and client is ", mongoDb)

}

func ConnectMongoDB(url string) (*mongo.Client, error) {
	if url == "" {
		return nil, errors.New("MONGO_URL is not set")
	}

	clientOptions := options.Client().ApplyURI(url)

	ctx := context.Background()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	return client, nil
}

func GetMongoDB() *mongo.Client {
	return mongoDb
}
