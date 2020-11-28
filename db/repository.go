package db

import (
	"time"

	"gorm.io/gorm"
)

// RepositoryInterface define interface for db repo
type RepositoryInterface interface {
	CreatePlayer(string) (*Player, error)
	CreatePlay(string) (Play, error)
	GetAllPlayers() []Player
	GetAllPlays() []Play
	GetPlayScoresForPlay(Play) []PlayScore
	GetPlayerByName(string) (Player, error)
	SaveScores(map[string]*SevenWondersScore) error
	SavePlayScore(Player, Play, SevenWondersScore) error
}

// Repository is a wrapper which encapsulate db access operation
type Repository struct {
	db *gorm.DB
}

// GetAllPlayers return all players from database
func (r *Repository) GetAllPlayers() []Player {
	var players []Player
	r.db.Find(&players)
	return players
}

// GetPlayScoresForPlay returns all play scores for play
func (r *Repository) GetPlayScoresForPlay(play Play) []PlayScore {
	var playScores []PlayScore
	r.db.Joins("Player").Joins("Play").Joins("SevenWondersScore").Where("play_id", play.ID).Find(&playScores)
	return playScores
}

// GetAllPlays returns all plays from database
func (r *Repository) GetAllPlays() []Play {
	var plays []Play
	r.db.Find(&plays)
	return plays
}

// CreatePlayer Create player with given name if that
// player does no exist
func (r *Repository) CreatePlayer(name string) (*Player, error) {
	player := Player{Name: name}
	result := r.db.Create(&player)
	return &player, result.Error
}

// GetPlayerByName returns player with given name or erorr
func (r *Repository) GetPlayerByName(name string) (Player, error) {
	player := Player{}
	res := r.db.Where("name = ?", name).First(&player)
	if err := res.Error; err != nil {
		return player, err
	}
	return player, nil
}

// CreatePlay saves play to database and returns it
func (r *Repository) CreatePlay(place string) (Play, error) {
	play := Play{Place: place, Date: time.Now()}
	res := r.db.Create(&play)
	if err := res.Error; err != nil {
		return play, err
	}
	return play, nil
}

// SavePlayScore saves given score, creates and saves playscore for it
func (r *Repository) SavePlayScore(player Player, play Play, score SevenWondersScore) error {
	res := r.db.Create(&score)
	if err := res.Error; err != nil {
		return err
	}
	playeScore := PlayScore{Player: player, Play: play, SevenWondersScore: score}
	res = r.db.Create(&playeScore)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

// SaveScores saves scores for given scores map
func (r *Repository) SaveScores(scoresMap map[string]*SevenWondersScore) error {
	play, err := r.CreatePlay("test")
	if err != nil {
		return err
	}
	for name, score := range scoresMap {
		player, err := r.GetPlayerByName(name)
		if err != nil {
			return err
		}
		err = r.SavePlayScore(player, play, *score)
		if err != nil {
			return err
		}
	}
	return nil
}
