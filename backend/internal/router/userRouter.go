package router

import "github.com/Allexsen/LoLGE-Wiki/backend/internal/handlers"

func setupUserRouter() {
	r.GET("/user/:id", handlers.UserGET)

	r.POST("/user/new", handlers.UserAdd)

	r.PUT("/user/edit/:id", handlers.UserEdit)

	r.DELETE("/user/delete/:id", handlers.UserDelete)

	r.GET("/user/:id/votes", handlers.UserVotesGET)
}
