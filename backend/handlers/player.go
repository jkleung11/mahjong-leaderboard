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
	identifier := c.Param("identifier")
	fmt.Printf("identifier is %v", identifier)
	var player models.Player

	if id, err := strconv.Atoi(identifier); err == nil {
		// number, so just query by primary key
		if err := h.DB.First(&player, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("id %v not found", identifier)})
			return
		}
	} else {
		if err := h.DB.Where("name = ?", identifier).First(&player).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("name %v not found", identifier)})
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
	if err := h.DB.Where("name = ?", request.NewName).First(&existing).Error; err == nil {
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
