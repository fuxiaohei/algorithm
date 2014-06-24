package trees

// a binary tree implement.
// use int value for example.


// binary tree node.
// it saves int value, pointers to left and right node.
type BinaryTreeNode struct {
	Value int
	LeftNode *BinaryTreeNode
	RightNode *BinaryTreeNode
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
			btn.LeftNode = NewBinaryTreeNode(i, nil, nil)
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
		btn.RightNode = NewBinaryTreeNode(i, nil, nil)
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
func NewBinaryTreeNode(i int, left *BinaryTreeNode, right *BinaryTreeNode) *BinaryTreeNode {
	return &BinaryTreeNode{
		Value:i,
		LeftNode:left,
		RightNode:right,
	}
}

// a comparison function to make sure the insert position.
// less one is in left or right.
type BinaryTreeCompareFn func(i, j int) bool

// binary tree struct
type BinaryTree struct {
	compareFn BinaryTreeCompareFn
	rootNode *BinaryTreeNode
}

// insert value into binary tree.
func (bt *BinaryTree) Insert(i int) {
	// if root node is missing, create this one for root.
	if bt.rootNode == nil {
		println("new root :", i)
		bt.rootNode = NewBinaryTreeNode(i, nil, nil)
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

// create new binary tree.
func NewBinaryTree(compareFn BinaryTreeCompareFn) *BinaryTree {
	b := new(BinaryTree)
	b.compareFn = compareFn
	return b
}

