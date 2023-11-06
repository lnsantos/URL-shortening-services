package presenter

import (
	"URLshortening/core"
	"URLshortening/domain"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type GetRegisterUrl struct {
	Endpoint string
}

func (c GetRegisterUrl) Create() (string, func(res http.ResponseWriter, req *http.Request)) {
	return c.Endpoint, func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		params := mux.Vars(req)
		short := params["short"]

		result, err := domain.GetUrl(short)

		fmt.Println("path short find is ", result)

		if err != nil {
			response, err := core.CreateErrorResponse(2, "failed in search short", err)
			res.WriteHeader(http.StatusBadRequest)

			if err == nil {
				_, _ = res.Write(response)
			} else {
				fmt.Println("unexpected err in send response: ", err)
			}

			return
		}

		response, _ := json.Marshal(result)
		res.WriteHeader(http.StatusOK)
		write, err := res.Write(response)
		fmt.Println(write)
		if err != nil {
			fmt.Println("unexpected err in send response: ", err)
		}
	}
}
