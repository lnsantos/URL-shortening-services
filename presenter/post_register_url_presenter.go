package presenter

import (
	"URLshortening/core"
	"URLshortening/domain"
	"encoding/json"
	"errors"
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
			} else {
				fmt.Println("unexpected err in send response: ", err)
			}
			return
		}

		var url string = input.Original

		short, err := domain.RegisterNewUrl(url)

		if err != nil {
			response, err := core.CreateErrorResponse(2, "failed in shorting url", err)
			res.WriteHeader(http.StatusBadRequest)

			if err == nil {
				_, _ = res.Write(response)
			} else {
				fmt.Println("unexpected err in send response: ", err)
			}

			return
		}

		if short == "" {
			response, err := core.CreateErrorResponse(3, "url already registered", errors.New(""))
			res.WriteHeader(http.StatusNotAcceptable)

			if err != nil {
				fmt.Println("unexpected err in send response: ", err)
				return
			}

			if _, err = res.Write(response); err != nil {
				fmt.Println("unexpected err in send response: ", err)
				return
			}

			return
		}

		responseData := RegisterUrlInput{Short: short, Original: url}
		res.WriteHeader(http.StatusCreated)

		if response, err := json.Marshal(responseData); err != nil {
			res.WriteHeader(http.StatusBadRequest)
			fmt.Println("unexpected err in send response: ", err)
		} else {
			if write, err := res.Write(response); err != nil {
				res.WriteHeader(http.StatusBadRequest)
				fmt.Println("unexpected err in send response: ", err)
			} else {
				fmt.Println("registered with success ", write)
			}
		}
	}
}
