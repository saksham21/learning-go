package main

import "testing"

func TestParadise(m *testing.T) {
	got := Paradise("Bali")
	want := "My paradise is Bali"
	if got != want {
		m.Errorf("Sum was incorrect, got: %s, want: %s.", got, want)
	}
}
