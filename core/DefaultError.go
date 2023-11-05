package core

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DefaultError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Debug       string `json:"debug,omitempty",`
}

func CreateErrorResponse(
	code int,
	description string,
	err error,
) ([]byte, error) {

	var hideDebug = true

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	} else {
		hideDebug = !(os.Getenv("ENV") == "dev")
	}

	var debugContent = ""
	if !hideDebug {
		debugContent = err.Error()
	}

	response := DefaultError{
		Code:        code,
		Description: description,
		Debug:       debugContent,
	}

	return json.Marshal(response)
}
