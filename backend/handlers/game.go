package handlers

import (
	"fmt"
	"mahjong-leaderboard-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GameHandler struct {
	DB *gorm.DB
}

type GameRequest struct {
	Date          string `json:"date" binding:"required"`
	WinnerID      *uint  `json:"winner_id" binding:"required"`
	WinningPoints *uint  `json:"winning_points" binding:"required"`
}

// create a game
func (h *GameHandler) CreateGame(c *gin.Context) {
	var request GameRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required fields in request"})
		return
	}

	parsedDate, err := time.Parse(time.RFC3339, request.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format, need RFC3339"})
		return
	}
	game := models.Game{
		Date:          parsedDate,
		WinnerID:      request.WinnerID,
		WinningPoints: request.WinningPoints,
	}

	// TODO: update ot check for valid winner id
	if err := h.DB.Create(&game).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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
