package server

import (
	"net/http"
)

func handleIndex(ctx *HandlerCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		plays := ctx.DB.GetAllPlays()
		var cards []Card

		for _, play := range plays {
			playScores := ctx.DB.GetPlayScoresForPlay(play)
			cards = append(cards, constructCardFromPlayScores(playScores))
		}

		ctx.Template.Render(w, cards)
	}
}

func handlePlayersList(ctx *HandlerCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		players := ctx.DB.GetAllPlayers()
		ctx.Template.Render(w, players)
	}
}
