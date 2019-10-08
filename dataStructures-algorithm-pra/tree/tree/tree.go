package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
