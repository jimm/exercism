package binary

func ParseBinary(binStr string) (val int, err error) {
	for i := 0; i < len(binStr); i++ {
		val *= 2
		if binStr[i] == '1' {
			val += 1
		}
	}
	return
}
