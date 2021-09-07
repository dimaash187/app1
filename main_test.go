package main

import (
	"testing"
)

func TestFactorialFunc(t *testing.T) {

	actual := 2
	if actual != Factorial(2) {
		t.Errorf("expected '2', got '%d'", actual)
	}

	actual = 6
	if actual != Factorial(3) {
		t.Errorf("expected '6', got '%d'", actual)
	}

	actual = 24
	if actual != Factorial(4) {
		t.Errorf("expected '24', got '%d'", actual)
	}

	actual = 120
	if actual != Factorial(5) {
		t.Errorf("expected '120', got '%d'", actual)
	}
}
