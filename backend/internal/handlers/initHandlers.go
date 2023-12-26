package handlers

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

func HomeGET() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.File(DIR_FRONT + "/html/home.html")
	}
}
