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
	r.GET("/", func(c *gin.Context) {
		c.File(DIR_FRONT + "/html/home.html")
	})

	setupFrontLinking()
	setupChampionRoutes()
	r.Run(":5000")
}

func setupFrontLinking() {
	r.GET("/html/:html/", func(c *gin.Context) {
		c.File(DIR_FRONT + "/html/" + c.Param("html"))
	})

	r.GET("/css/:styling/", func(c *gin.Context) {
		c.File(DIR_FRONT + "/css/" + c.Param("styling"))
	})

	r.GET("/js/:script/", func(c *gin.Context) {
		c.File(DIR_FRONT + "/js/" + c.Param("script"))
	})

	r.GET("/images/:image/", func(c *gin.Context) {
		c.File(DIR_FRONT + "/images/" + c.Param("image"))
	})
}

func setupChampionRoutes() {
	r.GET("/champions", handlers.ChampionsGET)
}
