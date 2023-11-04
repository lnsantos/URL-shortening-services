package domain

import (
	"URLshortening/data"
	"encoding/base64"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterNewUrl(
	c *mongo.Client,
	url string,
) (string, error) {

	encode := base64.StdEncoding.EncodeToString([]byte(url))
	record := data.RegisterUrlDTO{Original: url, Short: encode}

	err := data.RegisterUrl(c, record)

	return encode, err
}
