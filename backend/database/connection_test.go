package database

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestConnectToSupabase(t *testing.T) {
	// Load environment variables for testing
	err := godotenv.Load("../backend.env")
	if err != nil {
		t.Fatalf("Failed to load .env file: %v", err)
	}

	// Test the connection
	db, err := ConnectToSupabase()
	if err != nil {
		t.Fatalf("Failed to connect to Supabase: %v", err)
	}

	// Verify the connection by running a simple query
	// Use a unique query with timestamp to avoid prepared statement conflicts
	var result int
	err = db.Raw("SELECT 99 as test_connection_" + t.Name() + "_" + fmt.Sprintf("%d", time.Now().UnixNano())).Scan(&result).Error
	if err != nil {
		t.Fatalf("Failed to execute test query: %v", err)
	}

	if result != 99 {
		t.Errorf("Expected query result to be 99, got %d", result)
	}

	t.Log("✅ Successfully connected to Supabase and verified connection")
}

func TestConnectToSupabaseWithInvalidURL(t *testing.T) {
	// Test with missing environment variable
	originalURL := os.Getenv("SUPABASE_DATABASE_URL")
	defer os.Setenv("SUPABASE_DATABASE_URL", originalURL) // Restore original

	// Clear the environment variable
	os.Unsetenv("SUPABASE_DATABASE_URL")

	// Also clear any loaded .env variables by setting a dummy value
	os.Setenv("SUPABASE_DATABASE_URL", "")

	// Should fail
	_, err := ConnectToSupabase()
	if err == nil {
		t.Error("Expected error with missing SUPABASE_DATABASE_URL, but got nil")
	} else {
		t.Log("✅ Correctly handled missing environment variable")
	}
}
