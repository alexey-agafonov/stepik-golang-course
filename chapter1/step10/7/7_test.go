package main

import "testing"

const expectedYear = 8

func TestYearCount(t *testing.T) {
	year := YearCount(100, 10, 200)

	if year != expectedYear {
		t.Errorf("\nExpected: %v\nGot: %v", expectedYear, year)
	}
}
