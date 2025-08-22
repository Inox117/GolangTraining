package BinaryStringOperation

import (
	"strings"
)

// I found three different solutions.

// countOperationsToZeroSolution1
// Time complexity: O(n)
// Space complexity: O(1)
// This solution show the use of strings.Count and len.
func countOperationsToZeroSolution1(binary string) int {
	if binary == "0" {
		return 0
	}
	// if the bit is 1, we subtract 1 and divide by 2 -> 2 operations
	// if the bit is 0, we divide by 2 -> 1 operation
	numberOfZero := strings.Count(binary, "0")
	numberOfOne := len(binary) - numberOfZero
	// the last number is always 1, so we subtract 1 but we do not need to divide by 2
	numberOfOperations := numberOfZero + numberOfOne*2 - 1
	return numberOfOperations
}

// countOperationsToZeroSolution2
// Time complexity: O(n)
// Space complexity: O(1)
// This solution show the use of for loop and len, and does not use library functions.
func countOperationsToZeroSolution2(binary string) int {
	if binary == "0" {
		return 0
	}
	numberOfOperations := 0
	for i := len(binary) - 1; i >= 0; i-- {
		// if the bit is 1, we subtract 1 and divide by 2 -> 2 operations
		// if the bit is 0, we divide by 2 -> 1 operation
		if binary[i] == '1' {
			numberOfOperations += 2
		} else {
			numberOfOperations++
		}
	}
	// the last number is always 1, so we subtract 1 but we do not need to divide by 2
	return numberOfOperations - 1
}

// countOperationsToZeroSolution3
// Time complexity: O(n)
// Space complexity: O(n)
// This solution is about using recursion, which is a pain to maintain, debug and understand. Also, it is longer.
// I will never use this solution.
func countOperationsToZeroSolution3(binary string) int {
	if binary == "0" || binary == "" {
		return -1
	}
	lastBit := binary[0]
	if len(binary) > 1 {
		lastBit = binary[len(binary)-1]
	}
	if lastBit == '0' {
		// For an even number, we divide by 2
		// In binary, dividing by 2 is equivalent to right-shifting (removing the last '0')
		// However, we need to handle leading zeros correctly
		newBinary := countOperationsToZeroSolution3RemoveLeadingZero(binary[:len(binary)-1])
		if newBinary == "" {
			newBinary = "0"
		}
		return 1 + countOperationsToZeroSolution3(newBinary)
	} else {
		// For an odd number, we subtract 1 (which removes the last '1')
		// This is equivalent to removing the last digit in the binary string
		newBinary := binary[:len(binary)-1]
		if newBinary == "" {
			newBinary = "0"
		}
		return 2 + countOperationsToZeroSolution3(newBinary)
	}
}

func countOperationsToZeroSolution3RemoveLeadingZero(binary string) string {
	firstIndexOf1 := strings.Index(binary, "1")
	return binary[firstIndexOf1:]
}
