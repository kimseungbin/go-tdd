package iteration

// Repeat returns n times repeated string
func Repeat(input string, n int) (repeated string) {
	for i := 0; i < n; i++ {
		repeated += input
	}
	return repeated
}
