package main

import "testing"

const testResult = "I like Go!\nI like Go!\nI like Go!\n"

func TestTripleText(t *testing.T) {
	result := TripleText()

	if result != testResult {
		t.Errorf("\nExpected:\n%vGot:\n%v", testResult, result)
	}
}
