package router

import (
	"log"
	"os"

	"github.com/Allexsen/LoLGE-Wiki/backend/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	r *gin.Engine

	DIR_FRONT string
)

func GetEngine() *gin.Engine {
	return r
}

// Tester route
func InitRouter() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[Error]: Couldn't load .env file")
	}

	DIR_FRONT = os.Getenv("DIR_FRONT")
	r.GET("/", func(c *gin.Context) {
		c.File(DIR_FRONT + "/html/home.html")
	})

	r.Run(":5000")
}

func setupChampionRoutes() {
	r.GET("/champions/", handler.ChampionsGET)
}
