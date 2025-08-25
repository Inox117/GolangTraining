package StickFinder

import "slices"

func Solution(numbers []int, target int) []int {
	// We sort the array
	slices.Sort(numbers)
	// We iterate over the array
	for i := 0; i < len(numbers)-1; i++ {
		// For each number we check if the result of target - numbers[i] is present on the right side of the array
		// if yes, then we have found a valid pair, and because the array is sorted, it is also the solution with the shortest stick
		rest := target - numbers[i]
		if rest <= 0 {
			// if target is less than or equal to zero then we can stop.
			// It is not possible for a stick to have a negative length or a length equal to zero
			return []int{}
		}
		if slices.Contains(numbers[i+1:], rest) {
			return []int{numbers[i], rest}
		}
	}
	return []int{}
}
