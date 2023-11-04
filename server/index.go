package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func ServerStart(
	setting func(router *mux.Router),
) {

	router := mux.NewRouter()
	router.UseEncodedPath()
	custom := router.PathPrefix("/api").Subrouter()

	setting(custom)

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
