package router

import "github.com/Allexsen/LoLGE-Wiki/backend/internal/handlers"

func setupCommentRouter() {
	r.GET("/comments", handlers.CommentsGetAll)

	r.POST("/comments/:id", handlers.CommentAdd)

	r.PUT("/comments/:id", handlers.CommentEdit)

	r.DELETE("/comments/:id", handlers.CommentDelete)
}
