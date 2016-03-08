package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

var flowerNames = map[byte]string{
	'R': "radishes",
	'C': "clover",
	'G': "grass",
	'V': "violets",
}

var students = []string{
	"Alice", "Bob", "Charlie", "David", "Eve", "Fred",
	"Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry",
}

type Garden struct {
	students   []string
	kidFlowers map[string][]string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	if err := checkChildren(children); err != nil {
		return nil, err
	}
	lines, err := parseDiagram(diagram, len(children))
	if err != nil {
		return nil, err
	}

	g := Garden{kidFlowers: make(map[string][]string)}
	g.students = make([]string, len(children))
	copy(g.students, children)
	sort.Strings(g.students)
	for i := 0; i+1 < len(lines[0]); i += 2 {
		student := g.students[i/2]
		flowers := []string{
			flowerNames[lines[0][i]],
			flowerNames[lines[0][i+1]],
			flowerNames[lines[1][i]],
			flowerNames[lines[1][i+1]],
		}
		if flowers[0] == "" {
			return nil, errors.New("bad flower")
		}
		g.kidFlowers[student] = flowers
	}
	return &g, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	flowers := g.kidFlowers[child]
	if flowers == nil || flowers[0] == "" {
		return nil, false
	}
	return g.kidFlowers[child], true
}

func checkChildren(children []string) error {
	if len(children) == 0 {
		return errors.New("you forgot to think of the children")
	}
	for i := 0; i < len(children)-1; i++ {
		if children[i] == children[i+1] {
			return errors.New("clones not allowed")
		}
	}
	return nil
}

func parseDiagram(diagram string, numChildren int) ([]string, error) {
	lines := strings.Split(diagram, "\n")
	if lines[0] != "" {
		return nil, errors.New("bad garden format")
	}
	lines = lines[1:]
	if len(lines) < 2 || len(lines[0]) < numChildren || len(lines[1]) < numChildren {
		return nil, errors.New("garden isn't big enough")
	}
	if len(lines[0]) != len(lines[1]) {
		return nil, errors.New("garden isn't rectangular")
	}
	if (len(lines[0]) & 1) != 0 {
		return nil, errors.New("what an odd garden")
	}
	return lines, nil
}
