package store_test

import (
	"testing"
	"os"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == ""{
		databaseURL = "host=localhost dbname=planecontrol sslmode=disable user=postgres password=123"
	}

	os.Exit(m.Run())
}
