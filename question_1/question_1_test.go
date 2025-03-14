package main

import (
	"testing"
)

func TestFindMostValuePath(t *testing.T) {
	testCases := []struct {
		filePath string
		expected int
	}{
		{"hard.json", 7273},
		{"testcase.json", 237},
		{"empty1.json", -1},
		{"empty2.json", -1},
		{"no-file.json", -1},
	}
	folder := "./src/"
	for _, tc := range testCases {
		result := FindMostValuePath(folder + tc.filePath)
		if result != tc.expected {
			t.Errorf("FindMostValuePath(%s) = %d; want %d", tc.filePath, result, tc.expected)
		}
	}
}
