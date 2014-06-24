package main

import (
	"fmt"
	"github.com/fuxiaohei/algorithm/trees"
)

func main() {
	size := 30
	randomFind := trees.RandInt(1, size)
	randomValue := 0

	fn := func(i, j int) bool {
		return i < j
	}

	tree := trees.NewBinaryTree(fn)

	for i := 0; i < size; i++ {
		v := trees.RandInt(1, 999)
		tree.Insert(v)
		if i == randomFind {
			randomValue = v
		}
	}

	println("\n")

	tree.Display()

	println("\n")

	println("read in middle order:")
	fmt.Println(tree.ReadMiddle())

	println("\n")
	println("try to find value:", randomValue)
	tree.Find(randomValue)

}
