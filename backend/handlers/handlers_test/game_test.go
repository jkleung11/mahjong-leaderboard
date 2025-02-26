package handlers

import (
	"bytes"
	"encoding/json"
	"mahjong-leaderboard-backend/models"
	"mahjong-leaderboard-backend/services"
	"mahjong-leaderboard-backend/testutils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGame(t *testing.T) {

	testModels := []interface{}{
		&models.Player{},
		&models.GamePlayer{},
	}

	db, router, err := testutils.SetupTestEnvironment(testModels)
	if err != nil {
		t.Fatalf("Failed to set up test environment %v", err)
	}

	defer func() { db.Migrator().DropTable(testModels...) }()
	playerNames := []string{"leo", "raph", "don", "mich"}
	testutils.CreateTestPlayers(router, playerNames)

	gameData := map[string]interface{}{
		"date":           "2025-02-10T14:00:00Z",
		"winner":         "leo",
		"winning_points": 5,
		"players":        playerNames,
	}
	jsonBody, _ := json.Marshal(gameData)

	req, _ := http.NewRequest("POST", "/games", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code, "Expected 201 created")

	var responseStruct services.GameResponse
	err = json.Unmarshal(resp.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("failed to parse response body")
	}
	assert.NotZero(t, responseStruct.GameID, "Expected non-zero game ID")
	for _, player := range responseStruct.Players {
		if player.Name == "leo" {
			assert.Equal(t, "win", player.Result)
			assert.Equal(t, 5, int(player.PointsEarned))
			assert.Equal(t, "leo", player.Name)
		} else {
			assert.Equal(t, 0, int(player.PointsEarned))
			assert.Equal(t, player.Result, "loss")
		}
	}
}

// func TestCreateGameMissingField(t *testing.T) {
// 	// verify missing fields raise errors
// 	testModels := []interface{}{
// 		&models.Player{},
// 		&models.Game{},
// 	}

// 	_, router, err := testutils.SetupTestEnvironment(testModels)
// 	if err != nil {
// 		t.Fatalf("failed to set up test environment: %v", err)
// 	}

// 	missingField := map[string]interface{}{
// 		"date":      "2025-02-10T14:00:00Z",
// 		"winner_id": 1,
// 	}
// 	jsonBody, _ := json.Marshal(missingField)
// 	req := httptest.NewRequest("POST", "/games", bytes.NewBuffer(jsonBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	resp := httptest.NewRecorder()
// 	router.ServeHTTP(resp, req)
// 	assert.Equal(t, http.StatusBadRequest, resp.Code, "Expected bad request for missing field")
// }

// func TestCreateGameBadPlayer(t *testing.T) {
// 	testModels := []interface{}{
// 		&models.Player{},
// 		&models.Game{},
// 	}

// 	_, router, err := testutils.SetupTestEnvironment(testModels)
// 	if err != nil {
// 		t.Fatalf("failed to set up test environment: %v", err)
// 	}

// 	playerData := map[string]string{"name": "bob"}
// 	jsonBody, _ := json.Marshal(playerData)

// 	req := httptest.NewRequest("POST", "/players", bytes.NewBuffer(jsonBody))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp := httptest.NewRecorder()
// 	router.ServeHTTP(resp, req)

// 	gameData := map[string]interface{}{
// 		"date":      "2025-02-10T14:00:00Z",
// 		"winner_id": 2,
// 	}
// 	jsonBody, _ = json.Marshal(gameData)
// 	req = httptest.NewRequest("POST", "/games", bytes.NewBuffer(jsonBody))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp = httptest.NewRecorder()
// 	router.ServeHTTP(resp, req)
// 	assert.Equal(t, http.StatusBadRequest, resp.Code, "expect bad request for non existing winner")
// }
