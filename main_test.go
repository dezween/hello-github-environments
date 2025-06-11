package main

import "testing"

func TestGreet(t *testing.T) {
	expected := "Hello, GitHub!"
	if Greet("GitHub") != expected {
		t.Errorf("Expected %s, got %s", expected, Greet("GitHub"))
	}
}
