// +build !example

package tree

import (
	"errors"
)

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
				if r.Parent == c.ID {
					if r.ID < c.ID {
						return nil, errors.New("a")
					} else if r.ID == c.ID {
						if r.ID != 0 {
							return nil, errors.New("b")
						}
					} else {
						n++
						switch len(c.Children) {
						case 0:
							nn := &Node{ID: r.ID}
							c.Children = []*Node{nn}
							newTodo = append(newTodo, nn)
						case 1:
							nn := &Node{ID: r.ID}
							if c.Children[0].ID < r.ID {
								c.Children = []*Node{c.Children[0], nn}
							} else {
								c.Children = []*Node{nn, c.Children[0]}
							}
							newTodo = append(newTodo, nn)
						default:
							nn := &Node{ID: r.ID}
							newTodo = append(newTodo, nn)
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
		}
		todo = newTodo
	}
	if n != len(records) {
		return nil, errors.New("mismatch")
	}
	if err := chk(root, len(records)); err != nil {
		return nil, err
	}
	return root, nil
}

func chk(n *Node, m int) (err error) {
	if n.ID > m {
		return errors.New("z")
	} else if n.ID == m {
		return errors.New("y")
	} else {
		for i := 0; i < len(n.Children); i++ {
			err = chk(n.Children[i], m)
			if err != nil {
				return
			}
		}
		return
	}
}
