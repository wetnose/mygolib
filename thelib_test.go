package mygolib

import "testing"

func TestEcho(t *testing.T) {
	var s = echo("Hello")
	if s != "Hello" {
		t.Error("Expected 'Hello', got ", s)
	}
}
