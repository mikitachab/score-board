package server

import (
	"fmt"
	"net/http"

	"github.com/mikitachab/score-board/db"
)

func handlePlayersList(ctx HandlerCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var players []db.Player
		ctx.S.db.Find(&players)
		err := ctx.RenderTemplateFunc(w, players)
		handleErr(err)
	}
}

type AddPlayerViewCtx struct {
	Validated bool
	Errors    []string
}

func handleAddPlayer(ctx HandlerCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			err := ctx.RenderTemplateFunc(w, nil)
			handleErr(err)
		case "POST":
			r.ParseForm()
			playerName := r.FormValue("playerName")
			player := db.Player{Name: playerName}
			result := ctx.S.db.Create(&player)
			if result.Error != nil {
				errorMsg := fmt.Sprintf("Player with name %s already exist", playerName)
				err := ctx.RenderTemplateFunc(w, AddPlayerViewCtx{true, []string{errorMsg}})
				handleErr(err)
				return
			}
			http.Redirect(w, r, "/players", 302)
		}
	}
}
