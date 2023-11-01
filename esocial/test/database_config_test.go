package test

import (
	"os"
	"testing"
)

func TestDatabaseConfig(t *testing.T) {
	url := os.Getenv("MONGODB_URI")
	t.Log(url)
}
