package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(f func(int) bool) Ints {
	var retval Ints
	for _, val := range ints {
		if f(val) {
			retval = append(retval, val)
		}
	}
	return retval
}

func (ints Ints) Discard(f func(int) bool) Ints {
	var retval Ints
	for _, val := range ints {
		if !f(val) {
			retval = append(retval, val)
		}
	}
	return retval
}

func (lists Lists) Keep(f func([]int) bool) Lists {
	var retval Lists
	for _, val := range lists {
		if f(val) {
			retval = append(retval, val)
		}
	}
	return retval
}

func (strings Strings) Keep(f func(string) bool) Strings {
	var retval Strings
	for _, val := range strings {
		if f(val) {
			retval = append(retval, val)
		}
	}
	return retval
}
