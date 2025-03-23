package services

import (
	"errors"
	"mahjong-leaderboard-backend/dtos"
	"mahjong-leaderboard-backend/models"
	"time"

	"gorm.io/gorm"
)

func createGamePlayer(gameID uint, date time.Time, playerID uint, winnerID *uint, winningPoints *uint) models.GamePlayer {
	if winnerID == nil {
		return models.GamePlayer{GameID: gameID, Date: date, PlayerID: playerID, Result: string(models.Draw), PointsEarned: 0}
	}
	if playerID == *winnerID {
		return models.GamePlayer{GameID: gameID, Date: date, PlayerID: playerID, Result: string(models.Win), PointsEarned: *winningPoints}
	}
	return models.GamePlayer{GameID: gameID, Date: date, PlayerID: playerID, Result: string(models.Loss), PointsEarned: 0}
}

func CreateGamePlayers(gameID uint, date time.Time, players map[string]uint, winnerID *uint, winningPoints *uint) []models.GamePlayer {
	gamePlayers := make([]models.GamePlayer, 0, len(players))

	for _, playerID := range players {
		gamePlayer := createGamePlayer(gameID, date, playerID, winnerID, winningPoints)
		gamePlayers = append(gamePlayers, gamePlayer)
	}
	return gamePlayers

}

func FormatGameDetails(gameID uint, date time.Time, winner *string, players map[string]uint, gamePlayers []models.GamePlayer) dtos.GameDetails {
	playerIDToName := make(map[uint]string)
	for name, playerID := range players {
		playerIDToName[playerID] = name
	}

	var gameDetails dtos.GameDetails
	gameDetails.GameID = gameID
	gameDetails.Date = date
	gameDetails.Winner = winner

	for _, gamePlayer := range gamePlayers {
		name := playerIDToName[gamePlayer.PlayerID]
		playerResult := dtos.PlayerResults{Name: name, Result: gamePlayer.Result, PointsEarned: gamePlayer.PointsEarned}
		gameDetails.Results = append(gameDetails.Results, playerResult)
	}

	return gameDetails
}

func QueryGameDetailsByGameID(db *gorm.DB, gameID uint) (*dtos.GameDetails, error) {
	var gamePlayers []models.GamePlayer
	err := db.Where("game_id = ?", gameID).Preload("Player").Find(&gamePlayers).Error
	if err != nil {
		return nil, err
	}

	if len(gamePlayers) == 0 {
		return nil, errors.New("no game found for game id")
	}

	if len(gamePlayers) != 4 {
		return nil, errors.New("associated game id needs exactly 4 players")
	}
	date := gamePlayers[0].Date
	playerNameToID := map[string]uint{}
	var winner *string

	for _, gamePlayer := range gamePlayers {
		playerNameToID[gamePlayer.Player.Name] = gamePlayer.PlayerID
		if gamePlayer.Result == "win" {
			w := gamePlayer.Player.Name
			winner = &w
		}
	}
	gameDetails := FormatGameDetails(gameID, date, winner, playerNameToID, gamePlayers)
	return &gameDetails, nil

}
