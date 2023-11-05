package domain

import (
	"URLshortening/data"
	"go.mongodb.org/mongo-driver/bson"
)

type DataUrlOutput struct {
	Original string `json:"original"`
	Short    string `json:"short"`
}

func GetUrl(
	short string,
) (*DataUrlOutput, error) {

	filter := bson.D{{"short", short}}
	dto, err := data.SearchUrlData(filter)

	if err != nil {
		return nil, err
	}

	return &DataUrlOutput{Original: dto.Original, Short: dto.Short}, nil
}
