package handlers

import "github.com/gin-gonic/gin"

// Placeholder routes to define all initial paths

func VoteGET(c *gin.Context) {
	c.File(DIR_FRONT + "/html/vote")
}

func VotePOST(c *gin.Context) {
	c.File(DIR_FRONT + "/html/vote")
}
