package main

import (
	"github.com/fuxiaohei/algorithm/trees"
	"fmt"
)

func main() {
	size := 30
	randomFind := trees.RandInt(1, size)
	randomValue := int64(0)
	//tmpData := []int{308,31,158,559,83,17,46,488,501,577,195,226,478,98,928,79,446,689,120,130}


	tree := trees.NewAvlTree()

	for i := 0; i < size; i++ {
		v := int64(trees.RandInt(1, 999))
		//v := tmpData[i]
		tree.Insert(v,v+1)
		if i == randomFind {
			randomValue = v
		}
	}

	println("\n")
	tree.Dump(tree.Root,0,"root")

	println("read in asc order:")
	fmt.Println(tree.All(trees.AVL_ASC),len(tree.All(trees.AVL_ASC)))

	println("try to find value:", randomValue)
	fmt.Println(tree.Find(randomValue))

	tree.Remove(randomValue)
	println("after deleted, try above again.")

	tree.Dump(tree.Root,0,"root")

	println("read in asc order:")
	fmt.Println(tree.All(trees.AVL_ASC),len(tree.All(trees.AVL_ASC)))

	println("try to find value:", randomValue)
	fmt.Println(tree.Find(randomValue))
}
