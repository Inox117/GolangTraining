package HammingNumber

// Time complexity : O(n)
// Space complexity : O(n)

func Solution(n int) uint {
	// Array to store Hamming numbers
	hammingNumbers := make([]int, n)
	hammingNumbers[0] = 1
	nextHammingNumber := hammingNumbers[0]

	// The 3 index that we will use to browse the array
	i, j, k := 0, 0, 0

	// Next power for each
	nextPowerOf2 := 2
	nextPowerOf3 := 3
	nextPowerOf5 := 5

	for a := 1; a < n; a++ {
		// we get the minimal between the 3 nextPower
		nextHammingNumber = getMin(nextPowerOf2, getMin(nextPowerOf3, nextPowerOf5))
		hammingNumbers[a] = nextHammingNumber

		// We update the powers
		if nextHammingNumber == nextPowerOf2 {
			i = i + 1
			nextPowerOf2 = hammingNumbers[i] * 2
		}
		if nextHammingNumber == nextPowerOf3 {
			j = j + 1
			nextPowerOf3 = hammingNumbers[j] * 3
		}
		if nextHammingNumber == nextPowerOf5 {
			k = k + 1
			nextPowerOf5 = hammingNumbers[k] * 5
		}
	}
	return uint(nextHammingNumber)
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
