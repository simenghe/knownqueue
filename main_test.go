package main

import "testing"

func TestSquare(t *testing.T) {
	ans := Square(5)
	expected := 25
	if ans != expected {
		t.Errorf("Expected %d^2 = %d got %d", 5, expected, ans)
	}
}
