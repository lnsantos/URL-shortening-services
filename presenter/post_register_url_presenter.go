package presenter

import (
	"URLshortening/domain"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type PostRegisterUrl struct {
	Endpoint string
}

type RegisterUrlInput struct {
	Original string `json:"original"`
}

func (c PostRegisterUrl) Create(client *mongo.Client) (string, func(res http.ResponseWriter, req *http.Request)) {

	if client == nil {
		log.Fatal("client mongo is nil")
		return "", nil
	}

	return c.Endpoint, func(res http.ResponseWriter, req *http.Request) {

		var input RegisterUrlInput

		if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
			// customError := `{ "code": -1, "description": "failed in read body" }`
			res.WriteHeader(http.StatusBadRequest)
			_, _ = res.Write([]byte(err.Error()))
			return
		}
		fmt.Println(input.Original)

		var url string = input.Original

		short, errr := domain.RegisterNewUrl(client, url)

		fmt.Println("datsyaysgaysgaysgas ", short)

		if errr != nil {
			customError := `{ "code": 02, "description": "failed in shorting url" }`
			res.WriteHeader(http.StatusBadRequest)
			_, _ = res.Write([]byte(customError))
		}

		response := fmt.Sprintf(`{ "short": %v }`, short)
		res.WriteHeader(http.StatusOK)
		_, err := res.Write([]byte(response))

		if err != nil {
			fmt.Println("FINAL : ", err)
		}
	}
}
