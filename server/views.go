package server

import (
	"fmt"
	"net/http"
	"net/url"
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
			http.Redirect(w, r, "/players", http.StatusFound)
		}
	}
}

func handleAddScoreSelectPlayers(ctx *HandlerCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			players := ctx.DB.GetAllPlayers()
			ctx.Template.Render(w, players)
		case "POST":
			r.ParseForm()
			addScoreURL, err := url.Parse("/score/add/")
			if err != nil {
				panic(err)
			}
			q := addScoreURL.Query()
			for name := range r.Form {
				q.Add("player", name)
			}
			addScoreURL.RawQuery = q.Encode()
			http.Redirect(w, r, addScoreURL.String(), http.StatusFound)
		}
	}
}

func handleAddScore(ctx *HandlerCtx) http.HandlerFunc {
	categories := []string{
		"Military Conflicts",
		"Wonders",
		"Treasury Contents",
		"Civilian Structures",
		"Scientific Structures",
		"Commercial Structures",
		"Guilds",
	}
	type addScoreData struct {
		Players    []string
		Categories []string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			players := r.URL.Query()["player"]
			ctx.Template.Render(w, addScoreData{players, categories})
		case "POST":
			r.ParseForm()
			scores := GetScoresFromForm(r.Form)
			for k, v := range scores {
				fmt.Printf("%s %v", k, v)
			}
			ctx.DB.SaveScores(scores)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}
