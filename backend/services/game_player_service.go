package services

import (
	"mahjong-leaderboard-backend/models"
	"time"
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

func FormatGameResponse(gameID uint, players map[string]uint, gamePlayers []models.GamePlayer) GameResponse {
	playerIDToName := make(map[uint]string)
	for name, playerID := range players {
		playerIDToName[playerID] = name
	}

	var gameResponse GameResponse
	gameResponse.Message = "game created successfully"
	gameResponse.GameID = gameID

	for _, gamePlayer := range gamePlayers {
		name := playerIDToName[gamePlayer.PlayerID]
		playerResponse := PlayersResponse{Name: name, Result: gamePlayer.Result, PointsEarned: gamePlayer.PointsEarned}
		gameResponse.Players = append(gameResponse.Players, playerResponse)
	}

	return gameResponse
}
