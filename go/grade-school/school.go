package school

import (
	"fmt"
	"sort"
	"strings"
)

type Grade struct {
	grade int
	names []string
}

type School struct {
	grades map[int]Grade
}

func New() *School {
	return &School{grades: map[int]Grade{}}
}

func (s *School) Enrollment() []Grade {
	keys := []int{}
	for key := range s.grades {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	grades := []Grade{}
	for i, n := range keys {
		fmt.Println("Enrollment, i", i, "n", n)
		fmt.Println("s.grades[n].names", s.grades[n].names)
		names := s.grades[n].names
		grades[n].names = append(grades[n].names, strings.Join(names, " "))
	}
	return grades
}

func (s *School) Add(student string, grade int) {
	_, ok := s.grades[grade]
	if !ok {
		s.grades[grade] = Grade{grade: grade, names: []string{}}
	}
	fmt.Println("s.grades", s.grades)
	fmt.Println("s.grades[grade]", s.grades[grade])
	fmt.Println("s.grades[grade].names", s.grades[grade].names)
	s.grades[grade].names = append(s.grades[grade].names, student)
	sort.Strings(s.grades[grade].names)
}

func (s *School) Grade(grade int) []string {
	return s.grades[grade].names
}
