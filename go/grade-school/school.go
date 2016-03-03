package school

import "sort"

type Grade struct {
	grade int
	names []string
}

type School struct {
	grades map[int]*Grade
}

func New() *School {
	return &School{grades: map[int]*Grade{}}
}

func (s *School) Enrollment() []Grade {
	keys := []int{}
	for key := range s.grades {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	grades := []Grade{}
	for _, n := range keys {
		grades = append(grades, *s.grades[n])
	}
	return grades
}

func (s *School) Add(student string, grade int) {
	_, ok := s.grades[grade]
	if !ok {
		s.grades[grade] = &Grade{grade: grade, names: []string{}}
	}
	g := s.grades[grade]
	g.names = append(g.names, student)
	sort.Strings(g.names)
}

func (s *School) Grade(grade int) []string {
	g, ok := s.grades[grade]
	if ok {
		return g.names
	}
	return []string{}
}
