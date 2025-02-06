package handlers

import (
	"fmt"
	"mahjong-leaderboard-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PlayerHandler struct {
	DB *gorm.DB
}

// create
func (h *PlayerHandler) CreatePlayer(c *gin.Context) {
	var player models.Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing models.Player
	if err := h.DB.Where("name = ?", player.Name).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Player name already exists"})
	}

	if err := h.DB.Create(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create player"})
	}
	c.JSON(http.StatusCreated, player)
}

func (h *PlayerHandler) GetPlayer(c *gin.Context) {
	// check if id or name to get player
	query := c.Param("query")
	var player models.Player

	if _, err := strconv.Atoi(query); err != nil {
		// number, so just query by primary key
		if err := h.DB.First(&player, query).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("%v id not found", query)})
			return
		}
	} else {
		if err := h.DB.Where("name = ?").First(&player).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("%v name not found", query)})
			return
		}
	}
	c.JSON(http.StatusOK, player)
}

func (h *PlayerHandler) UpdatePlayer(c *gin.Context) {
	var request struct {
		CurrentName string `json:"current_name"`
		NewName     string `json:"new_name"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	var player models.Player
	if err := h.DB.Where("name = ?", request.CurrentName).First(&player).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("%v player not found", request.CurrentName)})
		return
	}

	var existing models.Player
	if err := h.DB.Where("name = ?", request.NewName).First(&existing).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("%v already exists", request.NewName)})
		return
	}

	player.Name = request.NewName

	if err := h.DB.Save(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update player"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player name updated", "player": player})

}
