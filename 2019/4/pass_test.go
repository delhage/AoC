package main

import (
	"testing"
)

func TestPass(t *testing.T) {

	tables := []struct {
		a string
		n bool
	}{
		{"111111",true},
		{"223450",false},
		{"123789",false},
	}

	for _, table := range tables {
		res := passOk(table.a)

		if res != table.n {
			t.Errorf("Password %s does not meet criteria.", table.a)
		}
	}
}
