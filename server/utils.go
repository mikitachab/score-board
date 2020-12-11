package server

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/mikitachab/score-board/db"
)

// GetScoresFromForm parse form values and returns scores
func GetScoresFromForm(form url.Values) map[string]*db.SevenWondersScore {
	data := flatForm(form)
	playersNames := getPlayersNames(data)
	delete(data, "players")
	scores := make(map[string]*db.SevenWondersScore)
	for _, name := range playersNames {
		if len(name) != 0 {
			scores[name] = &db.SevenWondersScore{}
		}
	}
	for key, value := range data {
		name, category := getNameAndCategory(key)
		if category == "Wonder" {
			scores[name].WonderName = value
		} else {
			score, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
			setScoreValue(scores[name], category, score)
		}
	}

	return scores
}

func getNameAndCategory(formInputName string) (string, string) {
	nameAndCategory := strings.Split(formInputName, "-")
	name, category := nameAndCategory[0], nameAndCategory[1]
	category = strings.ReplaceAll(category, " ", "")
	return name, category
}

func getPlayersNames(data map[string]string) []string {
	return strings.Split(data["players"], ";")
}

func flatForm(form url.Values) map[string]string {
	newMap := make(map[string]string)
	for k, v := range form {
		newMap[k] = v[0]
	}
	return newMap
}

func setScoreValue(s *db.SevenWondersScore, field string, value int) {
	switch field {
	case "MilitaryConflicts":
		s.MilitaryConflicts = value
	case "Wonders":
		s.Wonders = value
	case "TreasuryContents":
		s.TreasuryContents = value
	case "CivilianStructures":
		s.CivilianStructures = value
	case "ScientificStructures":
		s.ScientificStructures = value
	case "CommercialStructures":
		s.CommercialStructures = value
	case "Guilds":
		s.Guilds = value
	}
}
