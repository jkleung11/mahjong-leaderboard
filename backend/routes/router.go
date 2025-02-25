package routes

import (
	"mahjong-leaderboard-backend/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter sets up API routes
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	playerHandler := handlers.PlayerHandler{DB: db}
	gameHandler := handlers.GameHandler{DB: db}

	// Player routes
	r.GET("/players/:identifier", playerHandler.GetPlayer)
	r.POST("/players", playerHandler.CreatePlayer)
	r.PUT("/players", playerHandler.UpdatePlayer)

	// Game routes
	// r.GET("/games/:id", gameHandler.GetGameByID)
	r.POST("/games", gameHandler.CreateGame)

	return r
}
