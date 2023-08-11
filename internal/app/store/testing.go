package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseURL
	str := New(config)
	if err := str.Open(); err != nil {
		t.Fatal(err)
	}

	return str, func(tables ...string)  {
		if len(tables) > 0 {
			if _, err := str.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		str.Close()
	}
}