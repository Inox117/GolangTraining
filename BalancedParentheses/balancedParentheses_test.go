package BalancedParentheses

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		input    int
		expected []string
	}{
		{
			input:    0,
			expected: []string{""},
		},
		{
			input:    1,
			expected: []string{"()"},
		},
		{
			input:    2,
			expected: []string{"(())", "()()"},
		},
		{
			input:    3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}
	for _, tc := range testCases {
		result := Solution(tc.input)
		if len(result) != len(tc.expected) {
			t.Errorf("For %d, expected: %d, got: %d", tc.input, len(tc.expected), len(result))
		}
		for i := 0; i < len(result); i++ {
			if result[i] != tc.expected[i] {
				t.Errorf("For %d, expected: %s, got: %s", tc.input, tc.expected[i], result[i])
			}
		}
	}
}
