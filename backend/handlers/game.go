/*
Games from the front end will include all of our information (who played, who won).
This handler should validate that format includes what we need to create
entries for our Game and GamePlayers records.
*/

package handlers

import (
	"fmt"
	"mahjong-leaderboard-backend/models"
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

type GameHandler struct {
	DB *gorm.DB
}

type GameRequest struct {
	// Information provided from user on
	Date          string   `json:"date" binding:"required"`
	Winner        *string  `json:"winner" binding:"required"`
	WinningPoints *uint    `json:"winning_points" binding:"required"`
	Players       []string `json:"players" binding:"required"`
}

type GamePayload struct {
	// The actual payload for the record in the games table
	Date          string `json:"date" binding:"required"`
	WinnerID      *uint  `json:"winner_id" binding:"required"`
	WinningPoints *uint  `json:"winning_points" binding:"required"`
}

func ValidateGameRequest(request GameRequest) error {
	if len(request.Players) != 4 {
		return fmt.Errorf("A game must have exactly 4 players")
	}
	if len(request.Winner) > 0 {
		if slices.Contains(request.Players, &request.Winner) {
			return nil
		}
	}

	return nil
}

// create a game
func (h *GameHandler) CreateGame(c *gin.Context) {
	var payload GamePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required fields in payload"})
		return
	}

	parsedDate, err := time.Parse(time.RFC3339, payload.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format, need RFC3339"})
		return
	}
	game := models.Game{
		Date:          parsedDate,
		WinnerID:      payload.WinnerID,
		WinningPoints: payload.WinningPoints,
	}

	if err := h.DB.Create(&game).Error; err != nil {
		// parse db error specific to sqlite
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid winner_id, player does not exist"})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{"message": "game created successfully", "game": game})
}

// get a game based on id
func (h *GameHandler) GetGameByID(c *gin.Context) {
	id := c.Param("id")
	var game models.Game

	if err := h.DB.First(&game, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("game with id: %v not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusFound, game)
}
