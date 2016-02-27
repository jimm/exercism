package stringset

import (
	"fmt"
	"sort"
	"strings"
)

const testVersion = 3

type Set struct {
	set map[string]bool
}

func New() Set {
	return Set{set: make(map[string]bool)}
}

func NewFromSlice(strs []string) Set {
	set := New()
	for _, s := range strs {
		set.set[s] = true
	}
	return set
}

func (set Set) Add(s string) {
	set.set[s] = true
}

func (set Set) Delete(s string) {
	delete(set.set, s)
}

func (set Set) Has(s string) bool {
	return set.set[s]
}

func (set Set) IsEmpty() bool {
	return len(set.set) == 0
}

func (set Set) Len() int {
	return len(set.set)
}

func (set Set) Slice() []string {
	slice := make([]string, set.Len())
	i := 0
	for str := range set.set {
		slice[i] = str
		i++
	}
	return slice
}

func (set Set) String() string {
	strs := []string{}
	for str := range set.set {
		strs = append(strs, fmt.Sprintf("%#v", str))
	}
	return "{" + strings.Join(strs, ", ") + "}"
}

func Equal(s1, s2 Set) bool {
	if s1.Len() != s2.Len() {
		return false
	}
	sl1 := s1.Slice()
	sl2 := s2.Slice()
	sort.Strings(sl1)
	sort.Strings(sl2)
	for i := range sl1 {
		if sl1[i] != sl2[i] {
			return false
		}
	}
	return true
}

func Subset(s1, s2 Set) bool {
	if s1.Len() > s2.Len() {
		return false
	}
	for str := range s1.set {
		if !s2.Has(str) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for str := range s1.set {
		if s2.Has(str) {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	set := New()
	for str := range s1.set {
		if s2.Has(str) {
			set.set[str] = true
		}
	}
	for str := range s2.set {
		if s1.Has(str) {
			set.set[str] = true
		}
	}
	return set
}

func Union(s1, s2 Set) Set {
	set := New()
	for str := range s1.set {
		set.set[str] = true
	}
	for str := range s2.set {
		set.set[str] = true
	}
	return set
}

func Difference(s1, s2 Set) Set {
	set := New()
	for str := range s1.set {
		if !s2.Has(str) {
			set.set[str] = true
		}
	}
	return set
}

func SymmetricDifference(s1, s2 Set) Set {
	return Union(Difference(s1, s2), Difference(s2, s1))
}
