package storage

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseURL string) (*Storage, func(...string)) {
	// don`t test this method
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseURL

	s := New(config)

	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			query := fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
			if _, err := s.db.Exec(query); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
