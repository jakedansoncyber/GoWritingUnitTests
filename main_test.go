package main_test

import "testing"

// T object is injected by the testing framework itself
func TestAddition(t *testing.T) {
	got := 2 + 2
	expected := 4

	if got != expected {
		t.Errorf("Did not get expected result. Got : '%v', wanted '%v ", got, expected)
	}
}

func TestSubtraction(t *testing.T) {
	got := 10 - 5
	expected := 0
	if got != expected {
		t.Errorf("Did not get expected result. Got : '%v', wanted '%v ", got, expected)
	}
}
