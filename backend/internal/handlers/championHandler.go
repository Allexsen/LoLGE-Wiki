package handlers

import "github.com/gin-gonic/gin"

// Placeholder routes to define all initial paths

func ChampionsGetAll(c *gin.Context) {
	c.File(DIR_FRONT + "/html/champions")
}

func ChampionGET(c *gin.Context) {
	c.File(DIR_FRONT + "/html/champions/:champion")
}

func LoreGET(c *gin.Context) {
	c.File(DIR_FRONT + "/html/champions/:champion/lore")
}

func ChampionVotesGET(c *gin.Context) {
	// Should return the given votes
}

func ChampionVoteAdd(c *gin.Context) {
	// Should add a new vote
}

func ChampionVoteEdit(c *gin.Context) {
	// Should alter the given vote
}

func ChampionVoteDelete(c *gin.Context) {
	// Should delete the given vote
}
