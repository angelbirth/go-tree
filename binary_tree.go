package tree

import "golang.org/x/exp/constraints"

func NewBinaryTree[T constraints.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{
		isRoot: true,
	}
}

type BinaryTree[T constraints.Ordered] struct {
	left, right *BinaryTree[T]
	Val         T
	isRoot      bool
}

func (t *BinaryTree[T]) Search(x T) *BinaryTree[T] {
	if t.Val == x && !t.isRoot {
		return t
	}
	return t.searchChildren(x)
}

func (t *BinaryTree[T]) searchChildren(x T) *BinaryTree[T] {
	if t == nil {
		return nil
	}
	if t.Val == x {
		return t
	}
	if c := t.left.searchChildren(x); c != nil {
		return c
	}
	if c := t.right.searchChildren(x); c != nil {
		return c
	}
	return nil
}

func (t *BinaryTree[T]) AddChild(x T) *BinaryTree[T] {
	c := &BinaryTree[T]{
		Val: x,
	}
	if x > t.Val {
		t.right = c
	} else {
		t.left = c
	}
	return c
}

func (t *BinaryTree[T]) PreorderTraverse(f func(*BinaryTree[T])) {
	if !t.isRoot {
		f(t)
	}
	if t.left != nil {
		t.left.PreorderTraverse(f)
	}
	if t.right != nil {
		t.right.PreorderTraverse(f)
	}
}

func (t *BinaryTree[T]) PostorderTraverse(f func(*BinaryTree[T])) {
	if t.left != nil {
		t.left.PostorderTraverse(f)
	}
	if t.right != nil {
		t.right.PostorderTraverse(f)
	}
	if !t.isRoot {
		f(t)
	}
}
