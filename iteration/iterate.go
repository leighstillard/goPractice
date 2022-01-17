package iteration

func Repeat(character string, c int) string {
	var repeated string
	for i := 0; i < c; i++ {
		repeated += character
	}
	return repeated
}
