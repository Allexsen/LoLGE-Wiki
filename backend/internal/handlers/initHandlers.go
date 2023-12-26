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

	DIR_FRONT = os.Getenv("DIR_FRONT")
}

var (
	DIR_FRONT string // Frontend directory, asbolute path
)

func HomeGET(c *gin.Context) {
	c.File(DIR_FRONT + "/html/home.html")
}
