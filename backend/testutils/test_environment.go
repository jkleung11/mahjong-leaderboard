package testutils

import (
	"bytes"
	"encoding/json"
	"log"
	"mahjong-leaderboard-backend/handlers"
	"net/http"
	"net/http/httptest"

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
	router.POST("/games", gameHandler.CreateGame)
	router.GET("/games/:id", gameHandler.GetGameDetailsByID)

	return db, router, nil

}

func CreateTestPlayers(router *gin.Engine, playerNames []string) {
	for _, name := range playerNames {
		playerData := map[string]string{"name": name}
		jsonBody, _ := json.Marshal(playerData)

		req, _ := http.NewRequest("POST", "/players", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		log.Printf("created test player %v", name)
	}
}

func CleanupTestTables(db *gorm.DB, models []any) {
	for _, m := range models {
		db.Migrator().DropTable(m)
	}
}
