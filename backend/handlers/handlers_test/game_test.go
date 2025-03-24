package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mahjong-leaderboard-backend/dtos"
	"mahjong-leaderboard-backend/models"
	"mahjong-leaderboard-backend/testutils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

func TestCreateGame(t *testing.T) {

	testModels := []any{
		&models.Player{},
		&models.GamePlayer{},
	}

	db, router, err := testutils.SetupTestEnvironment(testModels)
	if err != nil {
		t.Fatalf("Failed to set up test environment %v", err)
	}

	playerNames := []string{"leo", "raph", "don", "mich"}
	testutils.CreateTestPlayers(router, playerNames)

	gameData := map[string]any{
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

	var responseStruct dtos.GameDetails
	err = json.Unmarshal(resp.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("failed to parse response body")
	}
	assert.NotZero(t, responseStruct.GameID, "Expected non-zero game ID")
	for _, playerResult := range responseStruct.Results {
		if playerResult.Name == "leo" {
			assert.Equal(t, "win", playerResult.Result)
			assert.Equal(t, 5, int(playerResult.PointsEarned))
			assert.Equal(t, "leo", playerResult.Name)
		} else {
			assert.Equal(t, 0, int(playerResult.PointsEarned))
			assert.Equal(t, playerResult.Result, "loss")
		}
	}
	defer testutils.CleanupTestTables(db, testModels)
}

func TestCreateGameMissingPlayer(t *testing.T) {
	testModels := []any{
		models.Player{},
		models.GamePlayer{},
	}
	db, router, err := testutils.SetupTestEnvironment(testModels)
	if err != nil {
		t.Fatalf("Failed to set up test environment %v", err)
	}

	testutils.CreateTestPlayers(router, []string{"leo", "raph", "don", "mich"})
	missingPlayer := []string{"leo", "raph", "don", "splinter"}
	gameData := map[string]any{
		"date":           "2025-02-10T14:00:00Z",
		"winner":         "leo",
		"winning_points": 5,
		"players":        missingPlayer,
	}
	jsonBody, _ := json.Marshal(gameData)

	req, _ := http.NewRequest("POST", "/games", bytes.NewBuffer(jsonBody))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code, "expected bad request due to unregistered player")

	defer testutils.CleanupTestTables(db, testModels)

}

func TestCreateGameMissingPoints(t *testing.T) {
	testModels := []any{
		models.Player{},
		models.GamePlayer{},
	}
	db, router, err := testutils.SetupTestEnvironment(testModels)
	if err != nil {
		t.Fatalf("Failed to set up test environment %v", err)
	}
	playerNames := []string{"leo", "raph", "don", "mich"}
	testutils.CreateTestPlayers(router, playerNames)
	gameData := map[string]any{
		"date":    "2025-02-10T14:00:00Z",
		"winner":  "leo",
		"players": playerNames,
	}
	jsonBody, _ := json.Marshal(gameData)
	req, _ := http.NewRequest("POST", "/games", bytes.NewBuffer(jsonBody))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	defer testutils.CleanupTestTables(db, testModels)
}

func TestGetGameByID(t *testing.T) {
	testModels := []any{
		models.Player{},
		models.GamePlayer{},
	}
	_, router, err := testutils.SetupTestEnvironment(testModels)
	if err != nil {
		t.Fatalf("Failed to set up test environment")
	}
	playerNames := []string{"leo", "raph", "don", "mich"}
	testutils.CreateTestPlayers(router, playerNames)
	gameData := map[string]any{
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

	// unpack the created game to compare with GET call
	var createdGameDetails dtos.GameDetails
	err = json.Unmarshal(resp.Body.Bytes(), &createdGameDetails)
	if err != nil {
		t.Fatalf("Failed to parse created response: %v", err)

	}
	assert.NotZero(t, createdGameDetails.GameID, "expected non zero game id")

	// compare get call with same id
	req, _ = http.NewRequest("GET", fmt.Sprintf("/games/%d", createdGameDetails.GameID), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "expected 200 for get game details")

	var retrievedGameDetails dtos.GameDetails
	err = json.Unmarshal(resp.Body.Bytes(), &retrievedGameDetails)
	if err != nil {
		t.Fatalf("Failed to parese GET game response: %v", err)
	}
	assert.Equal(t, createdGameDetails.GameID, retrievedGameDetails.GameID, "game id mismatch")
	assert.Equal(t, createdGameDetails.Date, retrievedGameDetails.Date, "date mismatch")
	assert.Equal(t, createdGameDetails.Winner, retrievedGameDetails.Winner, "winner mismatch")
	assert.Equal(t, len(playerNames), len(retrievedGameDetails.Results), "game id mismatch")

}
