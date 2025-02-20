package handlers

import (
	"bytes"
	"encoding/json"
	"mahjong-leaderboard-backend/models"
	"mahjong-leaderboard-backend/testutils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GameResponse struct {
	Message string      `json:"message"`
	Game    models.Game `json:"game"`
}

func TestCreateGame(t *testing.T) {

	testModels := []interface{}{
		&models.Game{},
		&models.Player{},
	}

	db, router, err := testutils.SetupTestEnvironment(testModels)
	if err != nil {
		t.Fatalf("Failed to set up test environment %v", err)
	}

	defer func() { db.Migrator().DropTable(testModels...) }()

	gameData := map[string]interface{}{
		"date":           "2025-02-10T14:00:00Z",
		"winner_id":      1,
		"winning_points": 3,
	}
	jsonBody, _ := json.Marshal(gameData)

	req, _ := http.NewRequest("POST", "/games", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusCreated, resp.Code, "Expected 201 created")

	var responseStruct GameResponse
	err = json.Unmarshal(resp.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("failed to parse response body")
	}

	assert.NotZero(t, responseStruct.Game.ID, "Expected non-zero game ID")

	assert.NotZero(t, responseStruct.Game.ID, "Expected non-zero game ID")
	assert.Equal(t, gameData["date"], responseStruct.Game.Date.Format("2006-01-02T15:04:05Z"), "Date mismatch")
	assert.Equal(t, uint(gameData["winner_id"].(int)), *responseStruct.Game.WinnerID, "Winner ID mismatch")
	assert.Equal(t, uint(gameData["winning_points"].(int)), *responseStruct.Game.WinningPoints, "Winning Points mismatch")

	// query the db to verify as well
	var dbGame models.Game
	if err := db.First(&dbGame, responseStruct.Game.ID).Error; err != nil {
		t.Fatalf("Game not found in db: %v", err)
	}
	assert.Equal(t, responseStruct.Game.ID, dbGame.ID, "DB id mismatch")
	assert.Equal(t, responseStruct.Game.Date, dbGame.Date, "DB date mismatch")
	assert.Equal(t, responseStruct.Game.WinnerID, dbGame.WinnerID, "DB winner id mismatch")
	assert.Equal(t, responseStruct.Game.WinningPoints, dbGame.WinningPoints, "DB winning points mismatch")

	// verify missing fields raise errors
	missingField := map[string]interface{}{
		"date":      "2025-02-10T14:00:00Z",
		"winner_id": 1,
	}
	jsonBody, _ = json.Marshal(missingField)
	req = httptest.NewRequest("POST", "/games", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code, "Expected bad request for missing field")

}
