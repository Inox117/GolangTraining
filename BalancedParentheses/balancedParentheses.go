package BalancedParentheses

func Solution(n int) []string {
	if n > 0 {
		return addParentheses("", n, 0, 0)
	}
	return []string{""}
}

func addParentheses(str string, n, open, close int) []string {
	var output []string
	if n == close {
		output = append(output, str)
	}
	if open < n {
		res := addParentheses(str+"(", n, open+1, close)
		output = append(output, res...)
	}
	if close < open {
		res := addParentheses(str+")", n, open, close+1)
		output = append(output, res...)
	}
	return output
}
