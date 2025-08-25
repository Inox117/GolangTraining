package Decompose

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		input    int64
		expected []int64
	}{
		{
			input:    11,
			expected: []int64{1, 2, 4, 10},
		},
		{
			input:    50,
			expected: []int64{1, 3, 5, 8, 49},
		},
	}
	for _, tc := range testCases {
		result := Solution(tc.input)
		if len(result) != len(tc.expected) {
			t.Errorf("For %d, expected length: %d, got: %d", tc.input, len(tc.expected), len(result))
		}
		for i := 0; i < len(result); i++ {
			if result[i] != tc.expected[i] {
				t.Errorf("For %d, expected: %d, got: %d", tc.input, tc.expected[i], result[i])
			}
		}
	}
}
