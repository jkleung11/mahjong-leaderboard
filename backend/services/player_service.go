package services

import (
	"mahjong-leaderboard-backend/models"

	"gorm.io/gorm"
)

func QueryPlayerIDsByNames(db *gorm.DB, names []string) (map[string]uint, error) {
	var players []models.Player
	db = db.Debug()
	err := db.Where("name in ?", names).Find(&players).Error
	if err != nil {
		return nil, err
	}

	playerMap := make(map[string]uint)
	for _, player := range players {
		playerMap[player.Name] = player.ID
	}

	return playerMap, nil
}

func GetWinnerID(winner *string, players map[string]uint) *uint {
	// given a map of player names and ids, return the winner's id
	if winner == nil {
		return nil
	}
	id := players[*winner]
	return &id
}
