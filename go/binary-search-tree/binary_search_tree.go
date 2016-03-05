package binarysearchtree

type SearchTreeData struct {
	data        int
	left, right *SearchTreeData
}

func Bst(n int) *SearchTreeData {
	return &SearchTreeData{data: n}
}

func (b *SearchTreeData) Insert(n int) {
	if n > b.data {
		if b.right != nil {
			b.right.Insert(n)
		} else {
			b.right = Bst(n)
		}
	} else {
		if b.left != nil {
			b.left.Insert(n)
		} else {
			b.left = Bst(n)
		}
	}
}

func (b *SearchTreeData) MapInt(f func(int) int) []int {
	var ints = []int{}
	if b.left != nil {
		ints = append(ints, b.left.MapInt(f)...)
	}
	ints = append(ints, f(b.data))
	if b.right != nil {
		ints = append(ints, b.right.MapInt(f)...)
	}
	return ints
}

func (b *SearchTreeData) MapString(f func(int) string) []string {
	var strings = []string{}
	if b.left != nil {
		strings = append(strings, b.left.MapString(f)...)
	}
	strings = append(strings, f(b.data))
	if b.right != nil {
		strings = append(strings, b.right.MapString(f)...)
	}
	return strings
}
