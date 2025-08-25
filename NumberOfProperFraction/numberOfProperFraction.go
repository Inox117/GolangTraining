package NumberOfProperFraction

func Solution(n int) int {
	if n == 1 {
		return 0
	}
	res := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			res -= res / i
		}
	}
	if n > 1 {
		res -= res / n
	}
	return res
}
