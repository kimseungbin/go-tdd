package iteration

func Repeat(input string) (repeated string) {
	for i := 0; i < 5; i++ {
		repeated += input
	}
	return repeated
}
