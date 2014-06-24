package main

import (
	"fmt"
	"github.com/fuxiaohei/algorithm/trees"
)

func main() {
	size := 20
	randomFind := trees.RandInt(1, size)
	randomValue := 0
	//tmpData := []int{308,31,158,559,83,17,46,488,501,577,195,226,478,98,928,79,446,689,120,130}

	fn := func(i, j int) bool {
		return i < j
	}

	tree := trees.NewBinaryTree(fn)

	for i := 0; i < size; i++ {
		v := trees.RandInt(1, 999)
		//v := tmpData[i]
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
	//tree.Find(randomValue)

	println("\n")
	tree.Del(randomValue)
	println("after deleted, try above again.")

	println("\n")
	tree.Display()

	println("\n")
	println("read in middle order:")
	fmt.Println(tree.ReadMiddle(),len(tree.ReadMiddle()))

	println("\n")
	println("try to find value:", randomValue)
	//tree.Find(randomValue)
}
