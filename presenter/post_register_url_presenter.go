package presenter

import (
	"URLshortening/core"
	"URLshortening/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

type PostRegisterUrl struct {
	Endpoint string
}

type RegisterUrlInput struct {
	Original string `json:"original,omitempty"`
	Short    string `json:"short,omitempty"`
}

func (c PostRegisterUrl) Create() (string, func(res http.ResponseWriter, req *http.Request)) {
	return c.Endpoint, func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		var input RegisterUrlInput

		if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
			response, err := core.CreateErrorResponse(-1, "failed in read body", err)
			res.WriteHeader(http.StatusBadRequest)

			if err == nil {
				_, _ = res.Write(response)
			}

			return
		}

		var url string = input.Original

		short, err := domain.RegisterNewUrl(url)

		fmt.Println("datsyaysgaysgaysgas ", short)

		if err != nil {
			response, err := core.CreateErrorResponse(2, "failed in shorting url", err)
			res.WriteHeader(http.StatusBadRequest)

			if err == nil {
				_, _ = res.Write(response)
			}

			return
		}

		response := fmt.Sprintf(`{ "short": %v }`, short)
		res.WriteHeader(http.StatusOK)
		_, err = res.Write([]byte(response))

		if err != nil {
			fmt.Println("unexpected err : ", err)
		}
	}
}
