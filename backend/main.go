package main

import (
	"log"
	"lutebard-backend/database"
)

func main() {
	// Connect to Supabase using GORM
	db, err := database.ConnectToSupabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Test the connection
	if err := database.TestConnection(db); err != nil {
		log.Fatalf("Database connection test failed: %v", err)
	}

	log.Println("✅ Successfully connected to database with GORM!")
	log.Println("🚀 Lutebard backend is ready!")
}
