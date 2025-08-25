package Decompose

func Solution(n int64) []int64 {
	return DecomposeRecursive(n*n, n)
}

func DecomposeRecursive(i, j int64) []int64 {
	if i < 0 {
		return nil
	}
	if i == 0 {
		return []int64{}
	}
	for k := j - 1; k >= 0; k-- {
		subArray := DecomposeRecursive(i-k*k, k)
		if subArray != nil {
			return append(subArray, k)
		}
	}
	return nil
}
