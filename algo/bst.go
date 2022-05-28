package algo

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

//https://leetcode.cn/problems/validate-binary-search-tree/
//验证二叉搜索树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Result struct {
	max  int
	min  int
	isOk bool
}

func isValidBST(root *TreeNode) bool {
	return dfs(root).isOk
}

func dfs(root *TreeNode) (rst Result) {
	if root == nil {
		rst.max = math.MinInt64
		rst.min = math.MaxInt64
		rst.isOk = true
		return
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	if root.Val > left.max && root.Val < right.min && left.isOk && right.isOk {
		rst.isOk = true
	}
	rst.max = max(max(left.max, right.max), root.Val)
	rst.min = min(min(left.min, right.min), root.Val)
	return
}

//https://leetcode.cn/problems/insert-into-a-binary-search-tree/
//二叉搜索树中的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

//https://leetcode.cn/problems/delete-node-in-a-bst/
//删除二叉搜索树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val == key {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	}
	return root
}

type Rst struct {
	height int //最大高度
	valid  bool
}

//判断是否是平衡二叉树
func isBalanced(root *TreeNode) bool {
	return isBalanced_dfs(root).valid
}
func isBalanced_dfs(root *TreeNode) (rst Rst) {
	if root == nil {
		rst.valid = true
		rst.height = 0
		return
	}
	left := isBalanced_dfs(root.Left)
	right := isBalanced_dfs(root.Right)
	if left.valid && right.valid && abs(left.height, right.height) <= 1 {
		rst.valid = true
	}
	rst.height = max(left.height, right.height) + 1
	return
}
