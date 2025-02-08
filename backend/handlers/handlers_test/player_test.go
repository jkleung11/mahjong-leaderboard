package handlers

import (
	"bytes"
	"encoding/json"
	"mahjong-leaderboard-backend/handlers"
	"mahjong-leaderboard-backend/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createTestDB() (*gorm.DB, *gin.Engine) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Player{})

	router := gin.Default()
	playerHandler := handlers.PlayerHandler{DB: db}
	router.POST("/players", playerHandler.CreatePlayer)
	return db, router
}

func TestCreatePlayer(t *testing.T) {
	db, router := createTestDB()
	defer func() { db.Migrator().DropTable(&models.Player{}) }() // drop table after test
	playerData := map[string]string{"name": "bob"}
	jsonData, _ := json.Marshal(playerData)

	req, _ := http.NewRequest("POST", "/players", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code, "Expected 201 created")

	var player models.Player
	err := db.First(&player, "name = ?", "bob").Error
	assert.Nil(t, err, "Player should exist in database")
	assert.Equal(t, "bob", player.Name, "Player names should match")

}
