package NumberOfProperFraction

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{
			input:    1,
			expected: 0,
		},
		{
			input:    2,
			expected: 1,
		},
		{
			input:    5,
			expected: 4,
		},
		{
			input:    15,
			expected: 8,
		},
		{
			input:    25,
			expected: 20,
		},
	}
	for _, tc := range testCases {
		result := Solution(tc.input)
		if result != tc.expected {
			t.Errorf("For %d, expected: %d, got: %d", tc.input, tc.expected, result)
		}
	}
}
