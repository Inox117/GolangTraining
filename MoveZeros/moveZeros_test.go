package MoveZeros

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			input:    []int{0, 0, 0, 0, 0},
			expected: []int{0, 0, 0, 0, 0},
		},
		{
			input:    []int{0, 0, 0, 0, 1},
			expected: []int{1, 0, 0, 0, 0},
		},
	}
	for _, tc := range testCases {
		result := Solution(tc.input)
		for i := 0; i < len(result); i++ {
			if result[i] != tc.expected[i] {
				t.Errorf("For %v, expected: %v, got: %v", tc.input, tc.expected, result)
			}
		}
	}
}
