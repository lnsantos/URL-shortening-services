package data

import (
	"URLshortening/infra"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func SearchUrlData(
	filter bson.D,
) (RegisterUrlDTO, error) {

	client, err := infra.ClientConnect()

	if err != nil {
		return RegisterUrlDTO{}, err
	}

	defer infra.OnDisconnectMongo(context.TODO(), client)

	var result RegisterUrlDTO

	mongoClient := infra.MongoClient{Mc: client}
	collection := mongoClient.GetCollectionShort()
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return RegisterUrlDTO{}, err
	}

	return result, nil
}

func SearchShortByUrl(
	url string,
) (string, error) {

	filter := bson.D{{"original", url}}
	data, err := SearchUrlData(filter)

	if err != nil {
		return "", err
	}

	return data.Short, nil
}

func SearchUrlByShort(
	short string,
) (string, error) {

	filter := bson.D{{"short", short}}
	data, err := SearchUrlData(filter)

	if err != nil {
		return "", err
	}

	return data.Short, nil
}
