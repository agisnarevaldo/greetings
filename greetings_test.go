package greetings

import "testing"

func TestHEllo(t *testing.T) {
	expected := "Hello, Go!"
	result := Hello("Go")

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
