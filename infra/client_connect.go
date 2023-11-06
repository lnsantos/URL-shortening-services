package infra

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func OnDisconnectMongo(ctx context.Context, client *mongo.Client) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	} else {
		fmt.Println("Client mongo is closed")
	}
}

func ClientConnect() (*mongo.Client, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	var mongoUri = os.Getenv("MONGO_URI")

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))

	defer cancel()

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println("There was a problem connecting to your Atlas cluster. Check that the URI includes a valid username and password, and that your IP address has been added to the access list. Error: ")
		panic(err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

type MongoClient struct {
	Mc *mongo.Client
}

func (c *MongoClient) GetCollectionShort() *mongo.Collection {
	return c.Mc.Database("url-short").Collection("shorting")
}
