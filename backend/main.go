package main

import (
	"log"
	"mahjong-leaderboard-backend/db"
	"mahjong-leaderboard-backend/models"
	"mahjong-leaderboard-backend/routes"
	"os"

	"gorm.io/gorm"
)

func main() {
	dbPath := os.Getenv("SQLITE_DB_PATH")
	if dbPath == "" {
		dbPath = "./db/mahjong.db"
	}

	database := db.ConnectDB(dbPath)

	// Check if we are running migrations instead of starting the server
	if os.Getenv("RUN_MIGRATIONS") == "true" {
		log.Println("Running migrations...")
		if err := RunMigrations(database); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		log.Println("Migrations completed successfully!")
		return
	}

	log.Println("Server is starting on port 8080...")
	router := routes.SetupRouter(database)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// RunMigrations applies database migrations
func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&models.Player{}, &models.Game{})
}
