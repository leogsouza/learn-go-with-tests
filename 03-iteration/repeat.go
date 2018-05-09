package iteration

func Repeat(character string) string {
	const repeatCount = 5
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
