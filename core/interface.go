package core

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type RegisterRouter interface {
	Create(c *mongo.Client) (string, func(res http.ResponseWriter, req *http.Request))
}
