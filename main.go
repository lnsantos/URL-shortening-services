package main

import (
	"URLshortening/presenter"
	"URLshortening/server"
	"github.com/gorilla/mux"
)

func main() {
	server.ServerStart(func(router *mux.Router) {
		router.HandleFunc(
			presenter.PostRegisterUrl{Endpoint: "/register/url"}.Create(),
		).Methods("POST")
		router.HandleFunc(
			presenter.GetRegisterUrl{Endpoint: "/url/{short}"}.Create(),
		).Methods("GET")
	},
	)
}
