package trees

import "strings"

// a binary tree implement.
// use int value for example.

// binary tree node.
// it saves int value, pointers to left and right node.
type BinaryTreeNode struct {
	Value     int
	ParentNode *BinaryTreeNode
	LeftNode  *BinaryTreeNode
	RightNode *BinaryTreeNode
}

// display tree node with simple structure scene.
func (btn *BinaryTreeNode) Display(deep int, title string) {
	println(strings.Repeat("|  ", deep-1)+"+  "+title+" (", btn.Value, ")")
	if btn.LeftNode != nil {
		btn.LeftNode.Display(deep+1, "Left")
	}
	if btn.RightNode != nil {
		btn.RightNode.Display(deep+1, "Right")
	}
}

// find value in binary tree nodes.
func (btn *BinaryTreeNode) Find(i int, fn BinaryTreeCompareFn, deep int) *BinaryTreeNode {
	if btn.Value == i {
		println(strings.Repeat("|  ", deep-2)+"* Find (", btn.Value, ")")
		return btn
	} else {
		// print root element
		if deep == 2 {
			println("+ Root (", btn.Value, ")")
		}
	}
	if fn(i, btn.Value) {
		// it should be in left node, but no left, so it's not found
		if btn.LeftNode == nil {
			return nil
		}
		println(strings.Repeat("|  ", deep-1)+"+ Left (", btn.LeftNode.Value, ")")
		return btn.LeftNode.Find(i, fn, deep+1)
	}

	if btn.RightNode == nil {
		return nil
	}
	println(strings.Repeat("|  ", deep-1)+"+ Right (", btn.RightNode.Value, ")")
	return btn.RightNode.Find(i, fn, deep+1)
}

// insert value into binary tree node.
// if compared true by fn, insert to left node.
// else insert to right node.
func (btn *BinaryTreeNode) Insert(i int, fn BinaryTreeCompareFn) {
	// if compared true, insert to left node.
	if fn(i, btn.Value) {
		// if left node is nil, add new node.
		if btn.LeftNode == nil {
			println("new left node :", i)
			btn.LeftNode = NewBinaryTreeNode(i, btn, nil, nil)
			return
		}
		// make left node to insert.
		println("insert into left-node:", btn.LeftNode.Value)
		btn.LeftNode.Insert(i, fn)
		return
	}

	// insert to right node.
	if btn.RightNode == nil {
		println("new right node :", i)
		btn.RightNode = NewBinaryTreeNode(i, btn, nil, nil)
		return
	}
	println("insert into right-node:", btn.RightNode.Value)
	btn.RightNode.Insert(i, fn)
}

// middle reading: left, node-self, right.
// now I only code middle-reading.
// the left-reading and right reading means the different order of left, node-self and right.
func (btn *BinaryTreeNode) ReadMiddle() []int {
	data := make([]int, 0)
	if btn.LeftNode != nil {
		data = btn.LeftNode.ReadMiddle()
	}
	data = append(data, btn.Value)
	if btn.RightNode != nil {
		data = append(data, btn.RightNode.ReadMiddle()...)
	}
	return data
}

// create new binary tree node with value and children.
func NewBinaryTreeNode(i int, parent , left, right *BinaryTreeNode) *BinaryTreeNode {
	return &BinaryTreeNode{
		Value:     i,
		LeftNode:  left,
		RightNode: right,
		ParentNode:parent,
	}
}

// a comparison function to make sure the insert position.
// less one is in left or right.
type BinaryTreeCompareFn func(i, j int) bool

// binary tree struct
type BinaryTree struct {
	compareFn BinaryTreeCompareFn
	rootNode  *BinaryTreeNode
}

// insert value into binary tree.
func (bt *BinaryTree) Insert(i int) {
	// if root node is missing, create this one for root.
	if bt.rootNode == nil {
		println("new root :", i)
		bt.rootNode = NewBinaryTreeNode(i, nil, nil, nil)
		println("insert done ----------")
		return
	}

	// insert to root node.
	println("insert into root")
	bt.rootNode.Insert(i, bt.compareFn)
	println("insert done ----------")
}

// read all binary tree in middle order.
func (bt *BinaryTree) ReadMiddle() []int {
	if bt.rootNode == nil {
		return nil
	}
	return bt.rootNode.ReadMiddle()
}

// print binary tree.
func (bt *BinaryTree) Display() {
	if bt.rootNode == nil {
		println("no node in binary tree")
		return
	}
	bt.rootNode.Display(1, "Root")
}

// find value in binary tree.
func (bt *BinaryTree) Find(i int) {
	if bt.rootNode == nil {
		println("no node in binary tree")
		return
	}
	if bt.rootNode.Find(i, bt.compareFn, 2) == nil {
		println("can not find", i, "in binary tree")
	}
}

func (bt *BinaryTree) Del(i int) {
	// try to find node.
	node := bt.rootNode.Find(i, bt.compareFn, 2)
	if node == nil {
		println("can not find", i, "in binary tree")
	}
	// if there is no children in this node, delete directly.
	if node.LeftNode == nil && node.RightNode == nil {
		if node == node.ParentNode.LeftNode {
			node.ParentNode.LeftNode = nil
			return
		}
		if node == node.ParentNode.RightNode {
			node.ParentNode.RightNode = nil
			return
		}
	}

	// if there is left child in this node.
	// the left child must be on left of parent node.
	// move it up.
	if node.RightNode == nil && node.LeftNode != nil {
		if node == node.ParentNode.LeftNode {
			node.ParentNode.LeftNode = node.LeftNode
		}else {
			node.ParentNode.RightNode = node.LeftNode
		}
		node.LeftNode.ParentNode = node.ParentNode
		return
	}

	// same to only left child case.
	// the right child must be on the right of parent node.
	// move it up.
	if node.LeftNode == nil && node.RightNode != nil {
		if node == node.ParentNode.LeftNode {
			node.ParentNode.LeftNode = node.RightNode
		}else {
			node.ParentNode.RightNode = node.RightNode
		}
		node.RightNode.ParentNode = node.ParentNode
		return
	}

	// if left and right child are existing, try to find last left node in right child.
	// the last left one is the min value in right child but over left child value.
	// so it can replace deleting node. its left node is old left child value.
	leftLast := node.RightNode
	for {
		if leftLast.LeftNode != nil {
			leftLast = leftLast.LeftNode
		}else {
			break
		}
	}
	if node == node.ParentNode.LeftNode {
		node.ParentNode.LeftNode = leftLast

	}else {
		node.ParentNode.RightNode = leftLast
	}
	leftLast.ParentNode = node.ParentNode
	leftLast.LeftNode = node.LeftNode
}

// create new binary tree.
func NewBinaryTree(compareFn BinaryTreeCompareFn) *BinaryTree {
	b := new(BinaryTree)
	b.compareFn = compareFn
	return b
}
