package tree

func NewTree[T comparable]() *Tree[T] {
	return &Tree[T]{
		isRoot: true,
	}
}

type Tree[T comparable] struct {
	children []*Tree[T]
	Val      T
	isRoot   bool
}

func (t *Tree[T]) Search(x T) *Tree[T] {
	if t.Val == x && !t.isRoot {
		return t
	}
	return t.searchChildren(x)
}

func (t *Tree[T]) searchChildren(x T) *Tree[T] {
	for _, child := range t.children {
		if child.Val == x {
			return child
		} else if c := child.searchChildren(x); c != nil {
			return c
		}
	}
	return nil
}

func (t *Tree[T]) AddChild(x T) *Tree[T] {
	c := &Tree[T]{
		Val: x,
	}
	t.children = append(t.children, c)
	return c
}

func (t *Tree[T]) AddChildren(x ...T) []*Tree[T] {
	var res []*Tree[T]
	for _, n := range x {
		c := &Tree[T]{
			Val: n,
		}
		t.children = append(t.children, c)
		res = append(res, c)
	}
	return res
}

func (t *Tree[T]) PreorderTraverse(f func(*Tree[T])) {
	if !t.isRoot {
		f(t)
	}
	for _, child := range t.children {
		child.PreorderTraverse(f)
	}
}

func (t *Tree[T]) PostorderTraverse(f func(*Tree[T])) {
	for _, child := range t.children {
		child.PostorderTraverse(f)
	}
	if !t.isRoot {
		f(t)
	}
}
