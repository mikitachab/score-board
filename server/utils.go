package server

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/mikitachab/score-board/db"
)

func validateNumPlayers(list []string) error {
	listLen := len(list)
	switch {
	case listLen < 2:
		return fmt.Errorf("not enough players, min:2 have:%d", listLen)
	case listLen > 7:
		return fmt.Errorf("to many players, max:7 have:%d", listLen)
	default:
		return nil
	}
}

func checkForDuplicates(list []string) error {
	validationMap := make(map[string]int)
	for _, element := range list {
		validationMap[element] += 1
	}

	for _, v := range validationMap {
		if v != 1 {
			// there is a duplicate here
			return fmt.Errorf("found duplicate in map")
		}
	}

	return nil
}

func validateAllPlayersInMap(list []string, checkMap map[string]bool) error {
	var err error = nil

	for _, name := range list {
		_, ok := checkMap[name]
		if !ok {
			if err == nil {
				err = fmt.Errorf("name(s) not present in database: %s", name)
			} else {
				err = fmt.Errorf("%v,%s", err, name)
			}
		}
	}

	return err
}

func checkScoreValue(value int) error {
	return nil
}

// GetScoresFromForm parse form values and returns scores
func GetScoresFromForm(form url.Values, players map[string]bool) (map[string]*db.SevenWondersScore, error) {
	data := flatForm(form)

	playersNames := getPlayersNames(data)
	delete(data, "players")

	err := validateNumPlayers(playersNames)
	if err != nil {
		return nil, err
	}

	err = checkForDuplicates(playersNames)
	if err != nil {
		return nil, err
	}

	err = validateAllPlayersInMap(playersNames, players)
	if err != nil {
		return nil, err
	}

	scores := make(map[string]*db.SevenWondersScore)
	for i, name := range playersNames {
		if len(name) != 0 {
			scores[name] = &db.SevenWondersScore{}
		} else {
			return nil, fmt.Errorf("empty name at index:%d", i)
		}
	}

	for key, value := range data {
		name, category := getNameAndCategory(key)

		//TODO{antonskwr} check for wonders duplicates
		if category == "Wonder" {
			scores[name].WonderName = value
		} else {
			score, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("can't construct score value")
			}

			err = checkScoreValue(score)
			if err != nil {
				return nil, err
			}

			setScoreValue(scores[name], category, score)
		}
	}

	return scores, nil
}

func getNameAndCategory(formInputName string) (string, string) {
	nameAndCategory := strings.Split(formInputName, "-")
	name, category := nameAndCategory[0], nameAndCategory[1]
	category = strings.ReplaceAll(category, " ", "")
	return name, category
}

func getPlayersNames(data map[string]string) []string {
	allPlayersStr := strings.Trim(data["players"], ";")
	return strings.Split(allPlayersStr, ";")
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
