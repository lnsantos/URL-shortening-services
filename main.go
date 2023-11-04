package main

import (
	"URLshortening/infra"
	"URLshortening/presenter"
	"URLshortening/server"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	client, err := infra.ClientConnect()

	if err != nil {
		log.Fatal(err)
		return
	}

	server.ServerStart(func(router *mux.Router) {
		router.HandleFunc(
			presenter.PostRegisterUrl{Endpoint: "/register/url"}.Create(client),
		).Methods("POST")
	})
}
