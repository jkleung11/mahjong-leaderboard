package handlers

import (
	"fmt"
	"mahjong-leaderboard-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GameHandler struct {
	DB *gorm.DB
}

// create a game
func (h *GameHandler) CreateGame(c *gin.Context) {
	var game models.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Create(&game).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"messge": "game created successfully", "game": game})
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
