package main

import "testing"

func TestAdd(m *testing.T) {
	got := add(7, 5)
	want := 12
	if got != want {
		m.Errorf("Sum was incorrect, got: %v, want: %v.", got, want)
	}
}

func TestSubtract(m *testing.T) {
	got := subtract(7, 5)
	want := 2
	if got != want {
		m.Errorf("Subtract was incorrect, got: %v, want: %v.", got, want)
	}
}

func TestDoMath(m *testing.T) {
	got := doMath(7, 5, add)
	want := 12
	if got != want {
		m.Errorf("Sum was incorrect, got: %v, want: %v.", got, want)
	}

	got = doMath(7, 5, subtract)
	want = 2
	if got != want {
		m.Errorf("Subtract was incorrect, got: %v, want: %v.", got, want)
	}
}
