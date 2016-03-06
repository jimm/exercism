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

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	if err := checkIds(records); err != nil {
		return nil, err
	}

	root := &Node{}
	todo := []*Node{root}
	for {
		if len(todo) == 0 {
			return root, nil
		}
		newTodo := []*Node(nil)
		for _, c := range todo {
			childRecordIndexes := []int{}
			for cri, r := range records {
				if r.Parent != c.ID || r.ID == r.Parent {
					continue
				}
				nn := &Node{ID: r.ID}
				newTodo = append(newTodo, nn)
				insertChild(c, r, nn)
				childRecordIndexes = append(childRecordIndexes, cri)
			}
			records = removeUsedRecords(records, childRecordIndexes)
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

func insertChild(c *Node, r Record, nn *Node) {
	for i, cc := range c.Children {
		if cc.ID > r.ID {
			c.Children = append(c.Children, nil)
			copy(c.Children[i+1:], c.Children[i:])
			c.Children[i] = nn
			return
		}
	}
	// not found, appending at end
	c.Children = append(c.Children, nn)
}

func removeUsedRecords(records []Record, childRecordIndexes []int) []Record {
	lcri := len(childRecordIndexes)
	if lcri == 0 {
		return records
	}

	workingRecords := make([]Record, len(records) - lcri)
	wrIndex := 0
	criIndex := 0
	for i, r := range records {
		if criIndex < lcri && childRecordIndexes[criIndex] == i {
			criIndex++
		} else {
			workingRecords[wrIndex] = r
			wrIndex++
		}
	}
	return workingRecords
}
