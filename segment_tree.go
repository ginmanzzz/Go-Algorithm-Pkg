package algo

type Node struct {
	left  *Node
	right *Node
	sum   int
	lazy  int
}

type SegmentTree struct {
	root_  *Node
	start_ int
	end_   int
}

/*
@params:

	node: 	to be updated.
	[start, end]: The interval managed by node.
	[l, r]: Operated interval.
*/
func (tree SegmentTree) update(node *Node, start, end, l, r, val int) {
	if l <= start && r >= end {
		node.sum += (end - start + 1) * val
		node.lazy += val
		return
	}
	mid := (start + end) / 2
	tree.pushDown(node, mid-start+1, end-mid)
	if l <= mid {
		tree.update(node.left, start, mid, l, r, val)
	}
	if r > mid {
		tree.update(node.right, mid+1, end, l, r, val)
	}
	tree.pushUp(node)
}

func (tree SegmentTree) pushDown(node *Node, leftNum, rightNum int) {
	if node.left == nil {
		node.left = new(Node)
	}
	if node.right == nil {
		node.right = new(Node)
	}
	if node.lazy == 0 {
		return
	}
	node.left.sum += node.lazy * leftNum
	node.right.sum += node.lazy * rightNum
	node.left.lazy += node.lazy
	node.right.lazy += node.lazy
	node.lazy = 0
}

func (tree SegmentTree) pushUp(node *Node) {
	node.sum = node.left.sum + node.right.sum
}

func (tree SegmentTree) query(node *Node, start, end, l, r int) int {
	if l <= start && r >= end {
		return node.sum
	}
	mid := (start + end) / 2
	ans := 0
	tree.pushDown(node, mid-start+1, end-mid)
	if l <= mid {
		ans += tree.query(node.left, start, mid, l, r)
	}
	if r > mid {
		ans += tree.query(node.right, mid+1, end, l, r)
	}
	return ans
}

func (tree *SegmentTree) Query(l, r int) int {
	return tree.query(tree.root_, tree.start_, tree.end_, l, r)
}

func (tree *SegmentTree) Update(l, r, val int) {
	tree.update(tree.root_, tree.start_, tree.end_, l, r, val)
}

func NewSegmentTree(max int) *SegmentTree {
	return &SegmentTree{
		new(Node),
		0,
		max,
	}
}
