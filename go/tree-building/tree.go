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
	r Record
	used bool
}

// Idea: newTodo should remember how many so we don't need to realloc

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
	todo := []*Node{root}
	for len(todo) > 0 {
		newTodo := []*Node(nil)
		for _, c := range todo {
			for _, rec := range recs {
				if rec.used || rec.r.Parent != c.ID || rec.r.ID == rec.r.Parent {
					continue
				}
				nn := &Node{ID: rec.r.ID}
				newTodo = append(newTodo, nn)
				insertChild(c, rec.r.ID, nn)
				rec.used = true
			}
		}
		todo = newTodo
	}
	return root, nil
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
	for i, cc := range c.Children {
		if cc.ID > id {
			c.Children = append(c.Children, nil)
			copy(c.Children[i+1:], c.Children[i:])
			c.Children[i] = nn
			return
		}
	}
	// not found, appending at end
	c.Children = append(c.Children, nn)
}
