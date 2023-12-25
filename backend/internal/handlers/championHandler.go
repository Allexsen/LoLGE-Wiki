package handlers

import "github.com/gin-gonic/gin"

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
