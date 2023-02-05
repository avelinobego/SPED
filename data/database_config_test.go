package data_test

import (
	_ "begoinformatica/sped/data"
	"os"
	"testing"
)

func TestDatabaseConfig(t *testing.T) {
	url := os.Getenv("MONGODB_URI")
	t.Log(url)
}
