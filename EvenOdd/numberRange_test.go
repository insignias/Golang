package main

import "testing"

func TestGenerateRange(t *testing.T) {
	n := numberRange{}

	n = generateRange(10)

	if len(n) != 11 {
		t.Errorf("Expected 11 number to be created, but got %v", len(n))
	}
}
