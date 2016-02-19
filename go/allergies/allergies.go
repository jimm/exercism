package allergies

var allergies = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(bits int) (achoo []string) {
	for i, s := range allergies {
		bit := 1 << uint(i)
		if bits&bit != 0 {
			achoo = append(achoo, s)
		}
	}
	return
}

func AllergicTo(bits int, allergy string) bool {
	allergicTo := Allergies(bits)
	for _, s := range allergicTo {
		if s == allergy {
			return true
		}
	}
	return false
}
