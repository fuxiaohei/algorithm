package trees

import "strings"

// AVL tree is a balanced binary tree with node deep condition and rotation.
// The deep means how many parent nodes are single-child nodes.
// If it's over deep, the node and parents node need rotate
// to keep node tree's height is not over deep.

const (
	AVL_ASC = iota + 1 // list avl tree as asc order
	AVL_DESC // list avl tree as desc order
)

// AvlTree node, saves key, value, balance factor
// and left,right, parent nodes.
type AvlTreeNode struct {
	Key     int64
	Value   interface{}
	Balance int
	Left, Right, Parent *AvlTreeNode
}

// find min node in node children.
func (n *AvlTreeNode) Min() *AvlTreeNode {
	if n.Left == nil {
		return n
	}
	min := n.Left
	for {
		if min.Left == nil {
			return min
		}
		min = min.Left
	}
	return min
}

// find max node in node children.
func (n *AvlTreeNode) Max() *AvlTreeNode {
	if n.Right == nil {
		return n
	}
	max := n.Right
	for {
		if max.Right == nil {
			return max
		}
		max = max.Right
	}
	return max
}

// list node children as asc order.
func (n *AvlTreeNode) Asc() []int64 {
	data := make([]int64, 0)
	if n.Left != nil {
		data = append(data, n.Left.Asc()...)
	}
	data = append(data, n.Key)
	if n.Right != nil {
		data = append(data, n.Right.Asc()...)
	}
	return data
}

// list node children as desc order.
func (n *AvlTreeNode) Desc() []int64 {
	data := make([]int64, 0)
	if n.Right != nil {
		data = append(data, n.Right.Desc()...)
	}
	data = append(data, n.Key)
	if n.Left != nil {
		data = append(data, n.Left.Desc()...)
	}
	return data
}

// find node by key in this node children.
func (n *AvlTreeNode) Find(key int64) *AvlTreeNode {
	if n.Key == key {
		return n
	}

	// try to find in left children
	if key < n.Key && n.Left != nil {
		return n.Left.Find(key)
	}

	// try to find in right children
	if key > n.Key && n.Right != nil {
		return n.Right.Find(key)
	}

	return nil
}

// get next node by key increment.
func (n *AvlTreeNode) Next() *AvlTreeNode {
	var next, self *AvlTreeNode

	if n.Right != nil {
		next = n.Right
		for {
			if next.Left == nil {
				break
			}
			next = next.Left
		}
		return next
	}

	next = n.Parent
	self = n
	for {
		if next != nil && self == next.Right {
			self = next
			next = self.Parent
			continue
		}
		break
	}
	return next
}

// get prev node by key decrement.
func (n *AvlTreeNode) Prev() *AvlTreeNode {
	var prev, self *AvlTreeNode
	if n.Left != nil {
		prev = n.Left
		for {
			if prev.Right == nil {
				break
			}
			prev = prev.Right
		}
		return prev
	}

	prev = n.Parent
	self = n
	for {
		if prev != nil && self == prev.Left {
			self = prev
			prev = self.Parent
			continue
		}
		break
	}
	return prev
}

func NewAvlTreeNode(k int64, v interface{}, p *AvlTreeNode) *AvlTreeNode {
	return &AvlTreeNode{Key:k, Value:v, Parent:p, Balance:0}
}

type AvlTree struct {
	Root *AvlTreeNode
}

func NewAvlTree() *AvlTree {
	return &AvlTree{}
}

func (tree *AvlTree) Iterator(asc int, fn func(*AvlTreeNode)) {
	if asc == AVL_ASC {
		next := tree.Root.Min()
		fn(next)
		for {
			next = next.Next()
			if next == nil {
				return
			}
			fn(next)
		}
		return
	}
	if asc == AVL_DESC {
		next := tree.Root.Max()
		fn(next)
		for {
			next = next.Prev()
			if next == nil {
				return
			}
			fn(next)
		}
	}
}

func (tree *AvlTree) Find(key int64) interface{} {
	if tree.Root != nil {
		res := tree.Root.Find(key)
		if res != nil {
			return res.Value
		}
	}
	return nil
}

func (tree *AvlTree) Dump(node *AvlTreeNode, level int, title string) {
	if node == nil {
		return
	}

	println(strings.Repeat("| ", level)+"+", title, "[", node.Key, "]", node.Balance)
	tree.Dump(node.Left, level+1, "left")
	tree.Dump(node.Right, level+1, "right")
}

func (tree *AvlTree) All(asc int) []int64 {
	if tree.Root == nil {
		return nil
	}
	if asc == AVL_ASC {
		return tree.Root.Asc()
	}
	if asc == AVL_DESC {
		return tree.Root.Desc()
	}
	return nil
}

func (tree *AvlTree) Insert(k int64, v interface{}) {
	if tree.Root == nil {
		tree.Root = NewAvlTreeNode(k, v, nil)
		return
	}
	tree.insertNode(tree.Root, NewAvlTreeNode(k, v, tree.Root))
}

func (tree *AvlTree) insertNode(p *AvlTreeNode, q *AvlTreeNode) {
	if q.Key < p.Key {
		if p.Left == nil {
			p.Left = q
			q.Parent = p
			avlRotate(p, tree)
			return
		}
		tree.insertNode(p.Left, q)
		return
	}
	if q.Key > p.Key {
		if p.Right == nil {
			p.Right = q
			q.Parent = p
			avlRotate(p, tree)
			return
		}
		tree.insertNode(p.Right, q)
		return
	}
	p.Value = q.Value
}

func (tree *AvlTree) Remove(k int64) {
	tree.removeNode(tree.Root, k)
}

func (tree *AvlTree) removeNode(n *AvlTreeNode, k int64) {
	if n == nil {
		return
	}
	if k < n.Key {
		tree.removeNode(n.Left, k)
		return
	}
	if k > n.Key {
		tree.removeNode(n.Right, k)
		return
	}
	if k == n.Key {
		tree.removeSelf(n)
	}
}

func (tree *AvlTree) removeSelf(q *AvlTreeNode) {
	var r *AvlTreeNode
	if q.Left == nil || q.Right == nil {
		if q.Parent == nil {
			tree.Root = nil
			q = nil
			return
		}
		r = q
	}else {
		r = q.Next()
		q.Key = r.Key
		q.Value = r.Value
	}

	var p *AvlTreeNode
	if r.Left != nil {
		p = r.Left
	}else {
		p = r.Right
	}

	if p != nil {
		p.Parent = r.Parent
	}

	if r.Parent == nil {
		tree.Root = p
	}else {
		if r.Parent.Left == r {
			r.Parent.Left = p
		}else {
			r.Parent.Right = p
		}
		avlRotate(r.Parent, tree)
	}

	r = nil
}

func avlHeight(n *AvlTreeNode) int {
	if n == nil {
		return -1
	}
	if n.Left == nil && n.Right == nil {
		return 0
	}
	l, r := avlHeight(n.Left), avlHeight(n.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}

func avlAdjustHeight(nodes ...*AvlTreeNode) {
	for _, n := range nodes {
		n.Balance = avlHeight(n.Right)-avlHeight(n.Left)
	}
}

func avlRotateLeft(n *AvlTreeNode) *AvlTreeNode {
	v := n.Right
	v.Parent = n.Parent
	n.Right = v.Left

	if n.Right != nil {
		n.Right.Parent = n
	}

	v.Left = n
	n.Parent = v

	if v.Parent != nil {
		if v.Parent.Right == n {
			v.Parent.Right = v
		}else {
			v.Parent.Left = v
		}
	}

	avlAdjustHeight(n, v)

	return v
}


func avlRotateRight(n *AvlTreeNode) *AvlTreeNode {
	v := n.Left
	v.Parent = n.Parent
	n.Left = v.Right

	if n.Left != nil {
		n.Left.Parent = n
	}

	v.Right = n
	n.Parent = v

	if v.Parent != nil {
		if v.Parent.Right == n {
			v.Parent.Right = v
		}else {
			v.Parent.Left = v
		}
	}

	avlAdjustHeight(n, v)

	return v
}


func avlRotateLeftRight(n *AvlTreeNode) *AvlTreeNode {
	n.Left = avlRotateLeft(n.Left)
	return avlRotateRight(n)
}

func avlRotateRightLeft(n *AvlTreeNode) *AvlTreeNode {
	n.Right = avlRotateRight(n.Right)
	return avlRotateLeft(n)
}

func avlRotate(n *AvlTreeNode, tree *AvlTree) {
	avlAdjustHeight(n)
	b := n.Balance

	if b == -2 {
		if avlHeight(n.Left.Left) >= avlHeight(n.Left.Right) {
			n = avlRotateRight(n)
		}else {
			n = avlRotateLeftRight(n)
		}
	}

	if b == 2 {
		if avlHeight(n.Right.Right) >= avlHeight(n.Right.Left) {
			n = avlRotateLeft(n)
		}else {
			n = avlRotateRightLeft(n)
		}
	}

	if n.Parent != nil {
		avlRotate(n.Parent, tree)
	}else {
		tree.Root = n
	}
}
