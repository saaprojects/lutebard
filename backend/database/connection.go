package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToSupabase() (*gorm.DB, error) {
	err := godotenv.Load("backend.env")
	if err != nil {
		err = godotenv.Load("../backend.env")
		if err != nil {
			log.Printf("Warning: Could not load backend.env file from either location: %v", err)
		}
	}

	databaseURL := os.Getenv("SUPABASE_DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("SUPABASE_DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db, nil
}

func TestConnection(db *gorm.DB) error {
	var result int
	// Use a unique query to avoid prepared statement conflicts
	err := db.Raw("SELECT 42 as test_value_" + fmt.Sprintf("%d", time.Now().UnixNano())).Scan(&result).Error
	if err != nil {
		return err
	}

	if result != 42 {
		log.Printf("Unexpected query result: expected 42, got %d", result)
	}

	return nil
}
