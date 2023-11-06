package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func ServerStart(
	setting func(router *mux.Router),
) {

	router := mux.NewRouter()
	router.UseEncodedPath()
	custom := router.PathPrefix("/api").Subrouter()

	setting(custom)
	middlewareLogger := handlers.LoggingHandler(os.Stdout, router)

	server := &http.Server{
		Handler:      middlewareLogger,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
