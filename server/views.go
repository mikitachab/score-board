package server

import (
	"fmt"
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

// AddPlayerFormCtx holds data
// for add player form tempalte
type AddPlayerFormCtx struct {
	Validated bool
	Errors    []string
}

func handleAddPlayer(ctx *HandlerCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			err := ctx.Template.Render(w, nil)
			handleErr(err)
		case "POST":
			r.ParseForm()
			playerName := r.FormValue("playerName")
			_, err := ctx.DB.CreatePlayer(playerName)
			if err != nil {
				errorMsg := fmt.Sprintf("Player with name %s already exist", playerName)
				viewCtx := AddPlayerFormCtx{true, []string{errorMsg}}
				ctx.Template.Render(w, viewCtx)
				return
			}
			http.Redirect(w, r, "/players", 302)
		}
	}
}
