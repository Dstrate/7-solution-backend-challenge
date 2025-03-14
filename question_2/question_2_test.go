package main

import (
	"testing"
)

func TestDecodeLeftRightEqual(t *testing.T) {
	testCases := []struct {
		encoded  string
		expected string
	}{
		{"LLRR=", "210122"},
		{"==RLL", "000210"},
		{"=LLRR", "221012"},
		{"RRL=R", "012001"},
		{"LRRLR", "101201"},
	}

	for _, tc := range testCases {
		result := DecodeLeftRightEqual(tc.encoded)
		if result != tc.expected {
			t.Errorf("DecodeLeftRightEqual(%s) = %s; want %s", tc.encoded, result, tc.expected)
		}
	}
}
