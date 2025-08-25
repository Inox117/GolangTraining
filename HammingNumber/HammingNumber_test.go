package HammingNumber

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		input    int
		expected uint
	}{
		{
			input:    1,
			expected: 1,
		},
		{
			input:    2,
			expected: 2,
		},
		{
			input:    3,
			expected: 3,
		},
		{
			input:    4,
			expected: 4,
		},
		{
			input:    5,
			expected: 5,
		},
		{
			input:    6,
			expected: 6,
		},
		{
			input:    7,
			expected: 8,
		},
		{
			input:    10,
			expected: 12,
		},
		{
			input:    12,
			expected: 16,
		},
		{
			input:    15,
			expected: 24,
		},
	}
	for _, tc := range testCases {
		result := Solution(tc.input)
		if result != tc.expected {
			t.Errorf("For %d, expected: %d, got: %d", tc.input, tc.expected, result)
		}
	}
}
