package handlers

import "github.com/gin-gonic/gin"

// Placeholder routes to define all initial paths

func UserGET(c *gin.Context) {
	// Should return the given user
}

func UserAdd(c *gin.Context) {
	// Should create a new user with given properties
}

func UserEdit(c *gin.Context) {
	// Should alter the existing user
}

func UserDelete(c *gin.Context) {
	// Should delete the given user
}

func UserVotesGET(c *gin.Context) {
	// Should return votes of the user
}
