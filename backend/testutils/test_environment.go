package testutils

import (
	"log"
	"mahjong-leaderboard-backend/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestEnvironment(models []interface{}) (*gorm.DB, *gin.Engine, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	log.Println("created test db in memory")

	if err := db.AutoMigrate(models...); err != nil {
		return nil, nil, err
	}

	log.Println("Ran migrations successfully")
	log.Println("Creating routes")
	router := gin.Default()
	// handlers
	playerHandler := handlers.PlayerHandler{DB: db}
	gameHandler := handlers.GameHandler{DB: db}
	// player routes
	router.GET("/players/:identifier", playerHandler.GetPlayer)
	router.POST("/players", playerHandler.CreatePlayer)
	router.PUT("/players", playerHandler.UpdatePlayer)
	// game routes
	router.GET("/games/:id", gameHandler.GetGameByID)
	router.POST("/games", gameHandler.CreateGame)

	return db, router, nil

}
