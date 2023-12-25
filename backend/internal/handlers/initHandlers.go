package handlers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[Error]: Couldn't load .env: " + err.Error())
	}
}

var (
	DIR_FRONT string = os.Getenv("DIR_FRONT") // Frontend directory, asbolute path
)
