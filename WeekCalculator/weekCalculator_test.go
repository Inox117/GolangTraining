package WeekCalculator

import (
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		input    []string
		expected int
	}{
		{
			input:    []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
			expected: 1,
		},
		{
			input:    []string{"Sun", "Sat", "Fri", "Thu", "Wed", "Tue", "Mon"},
			expected: 7,
		},
		{
			input:    []string{"Mon", "Mon", "Mon"},
			expected: 3,
		},
		{
			input:    []string{"Tue", "Sat", "Mon", "Fri"},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		result := Solution1(tc.input)
		if result != tc.expected {
			t.Errorf("For %s, expected: %d, got: %d", strings.Join(tc.input, ", "), tc.expected, result)
		}
		result = Solution2(tc.input)
		if result != tc.expected {
			t.Errorf("For %s, expected: %d, got: %d", strings.Join(tc.input, ", "), tc.expected, result)
		}
	}
}
