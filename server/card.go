package server

import "github.com/mikitachab/score-board/db"

// PlayerStat represents single stat
// in Histogram
type PlayerStat struct {
	Name  string
	Value int
}

// Histogram is used to display stats
type Histogram struct {
	Name          string
	Stats         []PlayerStat
	TopPlayerName string
}

// Card is used to represent information
// for single play on main page
type Card struct {
	Histograms []Histogram
}

func constructCardFromPlayScores(playScores []db.PlayScore) Card {
	histograms := []Histogram{
		Histogram{"Military", nil, "Not set"},
		Histogram{"Wonders", nil, "Not set"},
		Histogram{"Treasure", nil, "Not set"},
		Histogram{"Civil", nil, "Not set"},
		Histogram{"Science", nil, "Not set"},
		Histogram{"Commerce", nil, "Not set"},
		Histogram{"Guilds", nil, "Not set"},
	}

	for _, ps := range playScores {
		histograms[0].Stats = append(histograms[0].Stats, PlayerStat{ps.Player.Name, ps.SevenWondersScore.MilitaryConflicts})
		histograms[1].Stats = append(histograms[1].Stats, PlayerStat{ps.Player.Name, ps.SevenWondersScore.Wonders})
		histograms[2].Stats = append(histograms[2].Stats, PlayerStat{ps.Player.Name, ps.SevenWondersScore.TreasuryContents})
		histograms[3].Stats = append(histograms[3].Stats, PlayerStat{ps.Player.Name, ps.SevenWondersScore.CivilianStructures})
		histograms[4].Stats = append(histograms[4].Stats, PlayerStat{ps.Player.Name, ps.SevenWondersScore.ScientificStructures})
		histograms[5].Stats = append(histograms[5].Stats, PlayerStat{ps.Player.Name, ps.SevenWondersScore.CommercialStructures})
		histograms[6].Stats = append(histograms[6].Stats, PlayerStat{ps.Player.Name, ps.SevenWondersScore.Guilds})
	}

	for i, h := range histograms {
		histograms[i].TopPlayerName = findTopStat(h)
	}

	return Card{histograms}
}

func findTopStat(h Histogram) string {
	if len(h.Stats) == 0 {
		return "Empty"
	}

	topStat := 0
	topStatIndex := 0

	for i, stat := range h.Stats {
		if i == 0 || stat.Value > topStat {
			topStat = stat.Value
			topStatIndex = i
		}
	}

	return h.Stats[topStatIndex].Name
}
