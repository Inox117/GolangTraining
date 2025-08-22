package BinaryStringOperation

import (
	"strings"
	"testing"
)

func TestCountOperationsToZero(t *testing.T) {
	testCases := []struct {
		binary   string
		expected int
	}{
		{
			binary:   "1",
			expected: 1,
		},
		{
			binary:   "10",
			expected: 2,
		},
		{
			binary:   "11",
			expected: 3,
		},
		{
			binary:   "100",
			expected: 3,
		},
		{
			binary:   "101",
			expected: 4,
		},
		{
			binary:   strings.Repeat("1", 1000000), // 1 999 999 operations
			expected: 1999999,
		},
	}

	for _, tc := range testCases {
		result := countOperationsToZeroSolution1(tc.binary)
		if result != tc.expected {
			t.Errorf("Solution 1 binary %s, expected: %d, got: %d", tc.binary, tc.expected, result)
		}
		result = countOperationsToZeroSolution2(tc.binary)
		if result != tc.expected {
			t.Errorf("Solution 2 binary %s, expected: %d, got: %d", tc.binary, tc.expected, result)
		}
		result = countOperationsToZeroSolution3(tc.binary)
		if result != tc.expected {
			t.Errorf("Solution 3 binary %s, expected: %d, got: %d", tc.binary, tc.expected, result)
		}
	}
}
