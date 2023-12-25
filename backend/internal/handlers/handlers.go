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

// Placeholder routes to define all initial paths
func ChampionsGET(c *gin.Context) {
	c.File(DIR_FRONT + "/html/champions")
}

func ChampionGET(c *gin.Context) {
	c.File(DIR_FRONT + "/html/champions/:champion")
}

func LoreGET(c *gin.Context) {
	c.File(DIR_FRONT + "/html/champions/:champion/lore")
}

func VoteGET(c *gin.Context) {
	c.File(DIR_FRONT + "/html/vote")
}

func VotePOST(c *gin.Context) {
	c.File(DIR_FRONT + "/html/vote")
}

func CommentsGET(c *gin.Context) {
	// Should return JSON with comments' contents
}
