package storage_test

import (
	"os"
	"testing"
)

var databaseURL string

func TestMain(m *testing.M) {
	databaseURL = "user=postgres password=password host=localhost dbname=postgres sslmode=disable"

	os.Exit(m.Run())
}
