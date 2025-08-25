package StickFinder

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		inputStickLength    []int
		inputTargetedLength int
		expected            []int
	}{
		{
			inputStickLength:    []int{1, 2, 3, 4, 5},
			inputTargetedLength: 6,
			expected:            []int{1, 5},
		},
		{
			inputStickLength:    []int{1, 3, 3, 4},
			inputTargetedLength: 6,
			expected:            []int{3, 3},
		},
		{
			inputStickLength:    []int{1, 3, 4},
			inputTargetedLength: 6,
			expected:            []int{},
		},
		{
			inputStickLength:    []int{1},
			inputTargetedLength: 6,
			expected:            []int{},
		},
		{
			inputStickLength:    []int{1, 2, 9, 15},
			inputTargetedLength: 6,
			expected:            []int{},
		},
	}
	for _, tc := range testCases {
		result := Solution(tc.inputStickLength, tc.inputTargetedLength)
		if len(result) != len(tc.expected) {
			t.Errorf("Expected length %v, got %v", tc.expected, result)
		}
		for i, v := range result {
			if v != tc.expected[i] {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		}
	}
}
