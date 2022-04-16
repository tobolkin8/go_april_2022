package iteration

func Repeat(character string, repeatC int) string {
	var repeated string
	for i := 0; i < repeatC; i++ {
		repeated += character
	}
	return repeated
}
