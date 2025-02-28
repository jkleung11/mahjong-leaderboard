/*
Games from the front end will include all of our information (who played, who won).
This handler should validate that format includes what we need to create
entries for our Game and GamePlayers records.
*/

package handlers

import (
	"fmt"
	"mahjong-leaderboard-backend/models"
	"mahjong-leaderboard-backend/services"
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GameHandler struct {
	DB *gorm.DB
}

type GameRequest struct {
	// Information provided from user
	Date          string   `json:"date" binding:"required"`
	Winner        *string  `json:"winner" binding:"omitempty"`
	WinningPoints *uint    `json:"winning_points" binding:"omitempty"`
	Players       []string `json:"players" binding:"required"`
}

func validateGameRequest(request GameRequest) error {
	// check user input for who played in a game winning details
	if (request.Winner == nil) != (request.WinningPoints == nil) {
		return fmt.Errorf("winner and points must both be provided together")
	}
	if len(request.Players) != 4 {
		return fmt.Errorf("a game must have exactly 4 players")
	}
	if request.Winner != nil {
		if !slices.Contains(request.Players, *request.Winner) {
			return fmt.Errorf("winner must be one of the players in the game")
		}
	}
	return nil
}

func generateGameID(db *gorm.DB) uint {
	var latestGame models.GamePlayer
	db.Order("game_id DESC").First(&latestGame)
	return latestGame.GameID + 1
}

// create a game
func (h *GameHandler) CreateGame(c *gin.Context) {
	var request GameRequest
	// checks and early exits
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required fields in games request"})
		return
	}

	parsedDate, err := time.Parse(time.RFC3339, request.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format, need RFC3339"})
		return
	}

	if err := validateGameRequest(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// get the ids of players in the game
	players, err := services.QueryPlayerIDsByNames(h.DB, request.Players)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// a supplied player isn't registered
	if len(players) != 4 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "a game must have four registered players"})
		return
	}

	// create the structs needed for our insertions
	winnerID := services.GetWinnerID(request.Winner, players)
	gameID := generateGameID(h.DB)
	gamePlayers := services.CreateGamePlayers(gameID, parsedDate, players, winnerID, request.WinningPoints)

	transaction := h.DB.Begin()
	if err := transaction.Create(&gamePlayers).Error; err != nil {
		transaction.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	transaction.Commit()
	gameDetails := services.FormatGameDetails(gameID, parsedDate, request.Winner, players, gamePlayers)

	c.JSON(http.StatusCreated, gameDetails)
}
