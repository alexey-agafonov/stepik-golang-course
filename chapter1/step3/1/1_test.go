package main

import "testing"

const testResult = "I like Go!"

func TestText(t *testing.T) {
	result := Text()

	if result != testResult {
		t.Errorf("Expected \"%v\", got %v", testResult, result)
	}
}
