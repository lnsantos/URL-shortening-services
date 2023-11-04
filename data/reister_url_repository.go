package data

import (
	"URLshortening/infra"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type RegisterUrlDTO struct {
	Original string
	Short    string
}

func RegisterUrl(
	c *mongo.Client,
	record RegisterUrlDTO,
) error {

	client := infra.MongoClient{Mc: c}
	collection := client.GetCollectionShort()

	_, err := collection.InsertOne(
		context.TODO(),
		record,
	)

	if err != nil {
		fmt.Println("RegisterUrl:: ", err.Error())
		return err
	}

	return nil
}
