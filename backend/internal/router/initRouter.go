package router

import (
	"log"
	"os"

	"github.com/Allexsen/LoLGE-Wiki/backend/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	r = gin.Default()

	DIR_FRONT string
)

func GetEngine() *gin.Engine {
	return r
}

// Tester route
func InitRouter() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[Error]: Couldn't load .env file: " + err.Error())
	}

	DIR_FRONT = os.Getenv("DIR_FRONT")

	r.GET("/", handlers.HomeGET())
	r.GET("/home", handlers.HomeGET())

	setupChampionRouter()
	setupCommentRouter()
	setupFeDynamicLinkingRouter()
	setupUserRouter()

	r.Run(":5000")
}
