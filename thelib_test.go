package mygolib

import "testing"

func TestEcho(t *testing.T) {
	var s = Echo("Hello")
	if s != "Hello" {
		t.Error("Expected 'Hello', got ", s)
	}
}
