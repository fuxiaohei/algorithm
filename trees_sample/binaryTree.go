package main

import (
	"github.com/fuxiaohei/algorithm/trees"
	"fmt"
)

func main() {
	fn := func(i, j int) bool {
		return i < j
	}
	tree := trees.NewBinaryTree(fn)

	for i := 0; i < 30; i++ {
		tree.Insert(trees.RandInt(1, 9999))
	}

	fmt.Println(tree.ReadMiddle())
}

