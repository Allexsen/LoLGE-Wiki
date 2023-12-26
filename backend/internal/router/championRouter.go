package router

import "github.com/Allexsen/LoLGE-Wiki/backend/internal/handlers"

func setupChampionRouter() {
	r.GET("/champions", handlers.ChampionsGetAll)

	r.GET("/champions/:champion", handlers.ChampionGET)

	r.GET("/champions/:champion/lore", handlers.LoreGET)

	r.GET("/champions/:champion/votes", handlers.ChampionVotesGET)

	r.POST("/champions/:champion/votes/add", handlers.ChampionVoteAdd)

	r.PUT("/champions/:champion/votes/:id/edit", handlers.ChampionVoteEdit)

	r.DELETE("/champions/:champion/votes/:id/delete", handlers.ChampionVoteDelete)
}
