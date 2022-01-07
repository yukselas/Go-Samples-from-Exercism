package diffsquares

func SquareOfSum(n int) int {
	result := (n * (n + 1)) / 2
	return result * result
}

func SumOfSquares(n int) (sum int) {
	//	sum = (n * (n + 1) * (2*n + 1)) / 6
	for i := 0; i <= n; i++ {
		sum += i * i

	}
	return sum
}

func Difference(n int) (diff int) {
	diff = SquareOfSum(n) - SumOfSquares(n)
	return diff
}
