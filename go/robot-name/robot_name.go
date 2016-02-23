package robotname

import (
	"fmt"
	"math/rand"
)

type Robot struct {
	assignedName string
}

// Using a map means name lookups are O(1)
var names = map[string]bool{}

func (r *Robot) Name() string {
	if r.assignedName != "" {
		return r.assignedName
	}
	r.assignedName = genName()
	for names[r.assignedName] {
		r.assignedName = genName()
	}
	names[r.assignedName] = true
	return r.assignedName
}

func (r *Robot) Reset() {
	r.assignedName = ""
}

func genName() string {
	return fmt.Sprintf("%c%c%03d", randLetter(), randLetter(), rand.Int31n(1000))
}

func randLetter() byte {
	return byte(int32('A') + rand.Int31n(26))
}
