package main

import (
	"testing"
)

func TestProcessPassword(t *testing.T) {

	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Test with sample input",
			input:    `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`,
			expected: 6, // Adjust this value based on the expected result from your input file
		},
		{
			name:     "Test with empty input",
			input:    `L50
R101`,
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := processPassword(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}	