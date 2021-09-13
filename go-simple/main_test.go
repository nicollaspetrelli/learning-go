package main

import (
	"testing"
)

func TestAddition(t *testing.T) {
	if Addition(2, 2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}
