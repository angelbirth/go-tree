package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree_Search(t *testing.T) {
	var tests = []struct {
		desc  string
		init  func(*Tree[int])
		x     int
		found bool
	}{
		{
			desc:  "search empty tree",
			x:     1,
			found: false,
		},
		{
			desc: "search 1-child tree",
			x:    1,
			init: func(t *Tree[int]) {
				t.AddChild(1)
			},
			found: true,
		},
		{
			desc: "search in the middle of the tree",
			x:    4,
			init: func(t *Tree[int]) {
				t.AddChild(1).AddChild(2).AddChild(4).AddChild(8)
			},
			found: true,
		}, {
			desc: "search nonexisting value",
			x:    3,
			init: func(t *Tree[int]) {
				t.AddChild(1).AddChild(2).AddChild(4)
			},
			found: false,
		}, {
			desc: "search with >1 children, >1 depth",
			x:    5,
			init: func(t *Tree[int]) {
				node16 := t.AddChild(1).AddChild(2).AddChild(4).AddChild(8).AddChild(16)
				node16.AddChild(5)
				node16.AddChild(32)
			},
			found: true,
		},
		{
			desc: "search with >1 children, >1 depth, nonexisting value",
			x:    6,
			init: func(t *Tree[int]) {
				node16 := t.AddChild(1).AddChild(2).AddChild(4).AddChild(8).AddChild(16)
				node16.AddChild(5)
				node16.AddChild(32)
			},
			found: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tree := new(Tree[int])
			if tt.init != nil {
				tt.init(tree)
			}
			assert.Equal(t, tt.found, tree.Search(tt.x) != nil)
		})
	}

}

func TestTree_Traverse(t *testing.T) {
	tree := NewTree[rune]()
	f := tree.AddChild('F')
	b := f.AddChild('B')
	g := f.AddChild('G')
	b.AddChild('A')
	d := b.AddChild('D')
	d.AddChildren('C', 'E')
	g.AddChild('I').AddChild('H')
	tree.PreorderTraverse(func(t *Tree[rune]) {
		fmt.Printf("%c ", t.Val)
	})
	fmt.Println()
	tree.PostorderTraverse(func(t *Tree[rune]) {
		fmt.Printf("%c ", t.Val)
	})
}
