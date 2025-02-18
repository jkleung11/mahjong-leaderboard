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

func TestCreatePlayer(t *testing.T) {
	testModels := []interface{}{&models.Player{}}
	db, router, err := testutils.SetupTestEnvironment(testModels)
	if err != nil {
		t.Fatalf("Failed to set up test environment: %v", err)
	}

	defer func() { db.Migrator().DropTable(testModels...) }() // drop table after test

	playerData := map[string]string{"name": "bob"}
	jsonData, _ := json.Marshal(playerData)

	req, _ := http.NewRequest("POST", "/players", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code, "Expected 201 created")
}

func TestGetPlayerByName(t *testing.T) {
	testModels := []interface{}{&models.Player{}}
	db, router, err := testutils.SetupTestEnvironment(testModels)

	if err != nil {
		t.Fatalf("Failed to set up test environment: %v", err)
	}

	testPlayer := models.Player{
		Name: "bob",
	}

	db.Create(&testPlayer)
	defer func() { db.Migrator().DropTable(&models.Player{}) }() // drop table after test

	req := httptest.NewRequest("GET", "/players/bob", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected ok")

}
