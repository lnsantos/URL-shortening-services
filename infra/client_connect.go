package infra

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func onDisconnectMongo(ctx context.Context, client *mongo.Client) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func ClientConnect() (*mongo.Client, error) {
	var mongoUri = os.Getenv("MONGO_URI_URL_SHORTENING")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))

	defer cancel()
	defer onDisconnectMongo(ctx, client)

	err = client.Ping(ctx, nil)

	if err != nil {
		fmt.Println("There was a problem connecting to your Atlas cluster. Check that the URI includes a valid username and password, and that your IP address has been added to the access list. Error: ")
		panic(err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}
