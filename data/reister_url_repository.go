package data

import (
	"URLshortening/infra"
	"context"
	"fmt"
)

type RegisterUrlDTO struct {
	Original string
	Short    string
}

func RegisterUrl(
	record RegisterUrlDTO,
) error {

	client, err := infra.ClientConnect()

	defer infra.OnDisconnectMongo(context.TODO(), client)

	mongoClient := infra.MongoClient{Mc: client}
	collection := mongoClient.GetCollectionShort()

	_, err = collection.InsertOne(
		context.TODO(),
		record,
	)

	if err != nil {
		fmt.Println("RegisterUrl:: ", err.Error())
		return err
	}

	return nil
}
