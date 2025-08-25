package HumanReadableDuration

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		input    int64
		expected string
	}{
		{
			input:    1,
			expected: "1 second",
		},
		{
			input:    60,
			expected: "1 minute",
		},
		{
			input:    61,
			expected: "1 minute and 1 second",
		},
		{
			input:    3600,
			expected: "1 hour",
		},
		{
			input:    3661,
			expected: "1 hour, 1 minute and 1 second",
		},
		{
			input:    3600 * 24,
			expected: "1 day",
		},
		{
			input:    3600*24 + 3661,
			expected: "1 day, 1 hour, 1 minute and 1 second",
		},
		{
			input:    3600 * 24 * 365,
			expected: "1 year",
		},
		{
			input:    3600*24*365 + 3600*24 + 3661,
			expected: "1 year, 1 day, 1 hour, 1 minute and 1 second",
		},
	}
	for _, tc := range testCases {
		result := Solution(tc.input)
		if result != tc.expected {
			t.Errorf("For %d, expected: %s, got: %s", tc.input, tc.expected, result)
		}
	}
}
