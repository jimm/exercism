package pov

import "fmt"

const testVersion = 2

// A Graph is represented entirely by its arcs. The algorithm here is
// inefficient, mostly because parentOf needs to crawl through the entire
// list of arcs. Given the size of the graphs in the tests that doesn't
// matter. As an optimization of parentOf() we could store a map from
// children to parents once and update it within ChangeRoot. Or, you know,
// implement a real graph structure.

type Graph struct {
	arcs map[string][]string
}

func New() *Graph {
	return &Graph{arcs: map[string][]string{}}
}

func (g *Graph) AddNode(nodeLabel string) {
}

func (g *Graph) AddArc(from, to string) {
	g.arcs[from] = append(g.arcs[from], to)
}

func (g *Graph) ArcList() []string {
	list := []string{}
	for from, tos := range g.arcs {
		for _, to := range tos {
			list = append(list, fmt.Sprintf("%s -> %s", from, to))
		}
	}
	return list
}

func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	for parent := g.parentOf(newRoot); parent != ""; parent = g.parentOf(newRoot) {
		if grandparent := g.parentOf(parent); grandparent != "" {
			g.ChangeRoot(grandparent, parent)
		}
		g.reverseArc(parent, newRoot)
		parent = g.parentOf(newRoot)
	}
	return g
}

func (g *Graph) parentOf(label string) string {
	for from, tos := range g.arcs {
		for _, to := range tos {
			if to == label {
				return from
			}
		}
	}
	return ""
}

func (g *Graph) reverseArc(parent, child string) {
	g.AddArc(child, parent)
	g.removeArc(parent, child)
}

func (g *Graph) removeArc(parent, label string) {
	for i := 0; i < len(g.arcs[parent]); i++ {
		if g.arcs[parent][i] == label {
			// Is there a more efficient way to delete an element from an
			// array? Do I really have to allocate a new, smaller array?
			newChildren := make([]string, len(g.arcs[parent])-1)
			copy(newChildren, g.arcs[parent][0:i])
			copy(newChildren[i:], g.arcs[parent][i+1:])
			g.arcs[parent] = newChildren
			return
		}
	}
}
