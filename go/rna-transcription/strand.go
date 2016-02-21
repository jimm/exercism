package strand

const testVersion = 3

var xscription = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

func ToRNA(s string) string {
	retval := ""
	for _, b := range s {
		retval += xscription[b]
	}
	return retval
}
