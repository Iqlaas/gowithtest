package iteration

// const repeatCount = 5

// Repeat function
func Repeat(s string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += s
	}
	return repeated
}
