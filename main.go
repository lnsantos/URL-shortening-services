package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("variavel encontrada", os.Getenv("MONGO_URI_URL_SHORTENING"))
}
