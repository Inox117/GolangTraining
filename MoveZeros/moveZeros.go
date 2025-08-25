package MoveZeros

// Solution1 is the most straightforward solution.
// Time complexity: O(n)
// Space complexity: O(n)
func Solution1(arr []int) []int {
	toReturn := make([]int, len(arr))
	index := 0
	for _, value := range arr {
		if value != 0 {
			toReturn[index] = value
			index += 1
		}
	}
	return toReturn
}

// Solution2 is a more efficient solution.
// Time complexity: O(n)
// Space complexity: O(1)
// It is more efficient because we do not need to create a new array.
func Solution2(arr []int) []int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr[j] = arr[i]
			j++
		}
	}
	for k := j; k < len(arr); k++ {
		arr[k] = 0
	}
	return arr
}
