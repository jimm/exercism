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
	oneElementAnything := []int{1}
	n := 1
	for {
		if len(todo) == 0 {
			break
		}
		newTodo := []*Node(nil)
		for _, c := range todo {
			for _, r := range records {
				if r.Parent == c.ID && r.ID != r.Parent {
					nn := &Node{ID: r.ID}
					newTodo = append(newTodo, nn)
					n++
					switch len(c.Children) {
					case 0:
						c.Children = []*Node{nn}
					case 1:
						if c.Children[0].ID < r.ID {
							c.Children = []*Node{c.Children[0], nn}
						} else {
							c.Children = []*Node{nn, c.Children[0]}
						}
					default:
					breakpoint:
						for _ = range oneElementAnything {
							for _, cc := range c.Children {
								if cc.ID > r.ID {
									c.Children = append(c.Children, nil)
									copy(c.Children[1:], c.Children)
									c.Children[0] = nn
									break breakpoint

								}
							}
							c.Children = append(c.Children, nn)
						}
					}
				}
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
