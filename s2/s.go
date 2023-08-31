package s2

func Multiply(a, b int) int {
	x := 0
	if a == 0 || b == 0 {
		return 0
	}
	negCount := 0
	a, negCount = convertToPosAndOrIncrementNegCount(a, negCount)
	b, negCount = convertToPosAndOrIncrementNegCount(b, negCount)
	for i := 0; i < a; i++ {
		x += b
	}
	if negCount%2 == 0 {
		return x
	}
	return x * -1
}

func convertToPosAndOrIncrementNegCount(b int, negCount int) (int, int) {
	if b < 0 {
		b *= -1
		negCount++
	}
	return b, negCount
}
