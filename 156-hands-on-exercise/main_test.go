package main

import "testing"

func TestMain(m *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		m.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
