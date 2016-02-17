package pythagorean

type Triplet [3]int

func Range(min, max int) []Triplet {
	ts := make([]Triplet, 0)
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if a*a + b*b == c*c {
					ts = append(ts, Triplet{a, b, c})
				}
			}
		}
	}
	return ts
}

func Sum(p int) []Triplet {
	ts := make([]Triplet, 0)
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if a*a + b*b == c*c {
					ts = append(ts, Triplet{a, b, c})
				}
			}
		}
	}
	return ts

	// for _, t := range Range(1, p) {
	// 	if t[0] + t[1] + t[2] == p {
	// 		ts = append(ts, t)
	// 	}
	// }
	// return ts
}
