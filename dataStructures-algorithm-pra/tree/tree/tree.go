package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 相同的树
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

// 对称二叉树
func IsSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return (t1.Val == t2.Val) && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}

// 二叉树的最大深度
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	return max(left, right) + 1
}

// 二叉树中序遍历
func InorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := len(stack) - 1
		root = stack[n]
		stack = stack[:n]

		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

// 验证二叉搜索树
func IsValidBST(root *TreeNode) bool {
	var stack []*TreeNode

	min := -1 << 63
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := len(stack) - 1
		root := stack[n]
		stack = stack[:n]

		if root.Val <= min {
			return false
		}
		min = root.Val
		root = root.Right
	}
	return true
}
