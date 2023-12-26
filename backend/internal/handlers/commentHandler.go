package handlers

import "github.com/gin-gonic/gin"

// Placeholder routes to define all initial paths

func CommentsGetAll(c *gin.Context) {
	// Should return JSON with comments' contents
}

func CommentAdd(c *gin.Context) {
	// Should add the given comment to the db
}

func CommentDelete(c *gin.Context) {
	// Should remove the given comment from the db
}

func CommentEdit(c *gin.Context) {
	// Should edit the given comment in the db
}
