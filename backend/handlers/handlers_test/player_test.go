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
	testModels := []any{models.Player{}}
	db, router, err := testutils.SetupTestEnvironment(testModels)
	if err != nil {
		t.Fatalf("Failed to set up test environment: %v", err)
	}

	playerData := map[string]string{"name": "bob"}
	jsonBody, _ := json.Marshal(playerData)

	req, _ := http.NewRequest("POST", "/players", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code, "Expected 201 created")

	defer func() {
		if err := db.Migrator().DropTable(testModels...); err != nil {
			t.Fatalf("Error dropping tables after test: %v", err)
		}
	}()
}

func TestGetPlayerByName(t *testing.T) {
	testModels := []any{models.Player{}}
	db, router, err := testutils.SetupTestEnvironment(testModels)

	if err != nil {
		t.Fatalf("Failed to set up test environment: %v", err)
	}

	testPlayer := models.Player{
		Name: "bob",
	}

	db.Create(&testPlayer)

	req := httptest.NewRequest("GET", "/players/bob", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected ok")

	defer func() {
		if err := db.Migrator().DropTable(testModels...); err != nil {
			t.Fatalf("Error dropping tables after test: %v", err)
		}
	}()

}

func TestUpdatePlayerName(t *testing.T) {
	testModels := []any{models.Player{}}
	db, router, err := testutils.SetupTestEnvironment(testModels)

	if err != nil {
		t.Fatalf("Failed to set up test environment: %v", err)
	}

	testPlayer1 := models.Player{Name: "bob"}
	testPlayer2 := models.Player{Name: "linda"}

	db.Create(&testPlayer1)
	db.Create(&testPlayer2)

	// good update
	updateData := map[string]string{
		"current_name": "bob",
		"new_name":     "bobby",
	}
	jsonBody, _ := json.Marshal(updateData)

	req := httptest.NewRequest(http.MethodPut, "/players", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "expect ok code in response")

	// bad update, current name does not exist
	updateData = map[string]string{
		"current_name": "no name",
		"new_name":     "a name",
	}
	jsonBody, _ = json.Marshal(updateData)
	req = httptest.NewRequest("PUT", "/players", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code, "expect conflict code in response")
	// bad update, new name already exists
	updateData = map[string]string{
		"current_name": "bobby",
		"new_name":     "linda",
	}
	jsonBody, _ = json.Marshal(updateData)
	req = httptest.NewRequest("PUT", "/players", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusConflict, resp.Code, "expect conflict in response")

	defer func() {
		if err := db.Migrator().DropTable(testModels...); err != nil {
			t.Fatalf("Error dropping tables after test: %v", err)
		}
	}()

}
