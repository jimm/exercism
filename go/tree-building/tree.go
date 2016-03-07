package tree

import "errors"

const testVersion = 3

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

// Don't know if modifying Record is in the spirit of this exercise, so this
// is a wrapper.
type rec struct {
	r    Record
	used bool
}

// Pre-allocated list of nodes, so we don't have to keep reallocating.
type nodeList struct {
	nodes []*Node
	len   int
}

// Idea: both todo and newTodo should remember how many so we don't need to
// realloc

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	if err := checkIds(records); err != nil {
		return nil, err
	}

	recs := make([]rec, len(records))
	for i, r := range records {
		recs[i] = rec{r, false}
	}

	root := &Node{}
	todo := newNodeList(len(records))
	todo.add(root)
	newTodo := newNodeList(len(todo.nodes))
	for todo.len > 0 {
		createChildren(todo, &recs)
		copyNewChildren(todo, newTodo)

	}
	return root, nil
}

func createChildren(todo *nodeList, recs *[]rec) {
	for i := 0; i < todo.len; i++ {
		c := todo.nodes[i]
		for _, rec := range *recs {
			if rec.used || rec.r.Parent != c.ID || rec.r.ID == rec.r.Parent {
				continue
			}
			insertChild(c, rec.r.ID, &Node{ID: rec.r.ID})
			rec.used = true
		}
	}
}

// Copy all newly created children to todo list
func copyNewChildren(todo *nodeList, newTodo *nodeList) {
	numNewChildren := 0
	for i := 0; i < todo.len; i++ {
		copy(newTodo.nodes[numNewChildren:], todo.nodes[i].Children)
		numNewChildren += len(todo.nodes[i].Children)
	}
	copy(todo.nodes, newTodo.nodes)
	todo.len = numNewChildren
}

func checkIds(records []Record) error {
	numRecords := len(records)
	for _, r := range records {
		if r.ID < 0 || r.ID >= numRecords {
			return errors.New("illegal ID is out of range")
		} else if r.ID < r.Parent {
			return errors.New("child ID must be >= parent ID")
		} else if r.ID == r.Parent && r.ID != 0 {
			return errors.New("non-root child ID must not == parent ID")
		}
	}
	return nil
}

func insertChild(c *Node, id int, nn *Node) {
	c.Children = append(c.Children, nn)
	for i, cc := range c.Children {
		if cc.ID > id {
			copy(c.Children[i+1:], c.Children[i:])
			c.Children[i] = nn
			return
		}
	}
}

// ================ nodeList ================

func newNodeList(len int) *nodeList {
	return &nodeList{make([]*Node, len), 0}
}

func (todo *nodeList) add(nn *Node) {
	todo.nodes[todo.len] = nn
	todo.len++
}
