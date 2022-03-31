package main

import (
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历递归形式
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

//前序遍历非递归形式
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

//中序遍历递归形式
func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	preorderTraversal(root.Left)
	println(root.Val)
	preorderTraversal(root.Right)
}

//中序遍历非递归形式
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}

//后序遍历递归形式
func postorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
	println(root.Val)
}

//后序遍历非递归形式
func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var stack []*TreeNode
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}

//DFS 深度搜索-从上到下(前序遍历)
func preorderTraversal3(root *TreeNode) []int {
	var result []int
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

//DFS 深度搜索-从下到上(前序遍历)-分治法
func preorderTraversal4(root *TreeNode) []int {
	result := divideAndConquer(root)
	return result
}

func divideAndConquer(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

//归并排序
func MergeSort(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left []int, right []int) (result []int) {
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

//快速排序
func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start int, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quickSort(nums, 0, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

func partition(nums []int, start int, end int) int {
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	return i
}

//https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
//分治法 给定一个二叉树，找出其最大深度。
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

//https://leetcode-cn.com/problems/balanced-binary-tree/
//给定一个二叉树，判断它是否是高度平衡的二叉树。
//分治法 用-1 表示不平衡，>0 表示树高度
func isBalanced(root *TreeNode) bool {
	if maxDepth2(root) == -1 {
		return false
	}
	return true
}
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth2(root.Left)
	right := maxDepth2(root.Right)
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

//https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
//给定一个非空二叉树，返回其最大路径和。(Hard)
//思路：分治法，分为三种情况：
//左子树最大路径和最大，右子树最大路径和最大，左右子树最大加根节点最大，
//需要保存两个变量：一个保存子树最大路径和，一个保存左右加根节点和，然后比较这个两个变量选择最大值即可
type ResultType struct {
	SinglePath int // 保存单边最大值
	MaxPath    int // 保存最大值（单边或者两个单边+根的值）
}

func maxPathSum(root *TreeNode) int {
	result := helper(root)
	return result.MaxPath
}

func helper(root *TreeNode) ResultType {
	if root == nil {
		return ResultType{
			SinglePath: 0,
			MaxPath:    math.MinInt32,
		}
	}
	left := helper(root.Left)
	right := helper(root.Right)
	var result ResultType
	if left.SinglePath > right.SinglePath {
		result.SinglePath = max(left.SinglePath+root.Val, 0)
	} else {
		result.SinglePath = max(right.SinglePath+root.Val, 0)
	}
	maxPath := max(right.MaxPath, left.MaxPath)
	result.MaxPath = max(maxPath, left.SinglePath+right.SinglePath+root.Val)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
//给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

////----------BFS----------

//https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
//BFS层次遍历
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var list []int
		n := len(queue) //queue长度会变化
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, list)
	}
	return result
}

//https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
//给定一个二叉树，返回其节点值自底向上的层次遍历
func levelOrderBottom(root *TreeNode) [][]int {
	result := levelOrder(root)
	reverse(result)
	return result
}

func reverse(nums [][]int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

//https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
//给定一个二叉树，返回其节点值的锯齿形层次遍历。Z 字形遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var isZ bool
	for len(queue) > 0 {
		var list []int
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if isZ {
			reverse2(list)
		}
		result = append(result, list)
		isZ = !isZ
	}
	return result
}
func reverse2(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

///----------二叉搜索树----------
//https://leetcode-cn.com/problems/validate-binary-search-tree/
//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
func isValidBST(root *TreeNode) bool {
	var result []int
	inOrder(root, &result)
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}
func inOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, result)
	*result = append(*result, root.Val)
	inOrder(root.Right, result)
}

//https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/
//给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

//https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
//从中序与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	node := postorder[len(postorder)-1]
	index := func(inorder []int, node int) int {
		for i, v := range inorder {
			if v == node {
				return i
			}
		}
		return -1
	}(inorder, node)
	root := &TreeNode{
		Val:   node,
		Left:  buildTree(inorder[:index], postorder[:index]),
		Right: buildTree(inorder[index+1:], postorder[index:len(postorder)-1]),
	}
	return root
}

//https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
//从前序与中序遍历序列构造二叉树
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 || len(preorder) == 0 {
		return nil
	}
	node := preorder[0]
	index := func(inorder []int, node int) int {
		for i, v := range inorder {
			if v == node {
				return i
			}
		}
		return -1
	}(inorder, node)
	root := &TreeNode{
		Val:   node,
		Left:  buildTree(preorder[1:index+1], inorder[:index]),
		Right: buildTree(preorder[index+1:], inorder[index+1:]),
	}
	return root
}

//https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/
//填充每个节点的下一个右侧节点指针
type NodeN struct {
	Val   int
	Left  *NodeN
	Right *NodeN
	Next  *NodeN
}

func connect(root *NodeN) *NodeN {
	setNext(root, nil)
	return root
}
func setNext(root *NodeN, next *NodeN) {
	if root == nil {
		return
	}
	root.Next = next
	setNext(root.Left, root.Right)
	var rightNode *NodeN
	if root.Next != nil {
		rightNode = root.Next.Left
	}
	setNext(root.Right, rightNode)
}

//https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/
//填充每个节点的下一个右侧节点指针 II
func connect2(root *NodeN) *NodeN {
	setNext(root, nil)
	return root
}
func setNext2(root *NodeN, next *NodeN) {
	if root == nil {
		return
	}
	getNode := func(node *NodeN) (subNode *NodeN) {
		for node != nil {
			if node.Left != nil {
				subNode = node.Left
				break
			}
			if node.Right != nil {
				subNode = node.Right
				break
			}
			node = node.Next
		}
		return
	}
	root.Next = next
	setNext(root.Right, getNode(root.Next))
	setNext(root.Left, func() *NodeN {
		if root.Right != nil {
			return root.Right
		} else {
			return getNode(root.Next)
		}
	}())
}

//https://leetcode-cn.com/leetbook/read/data-structure-binary-tree/xomr73/
//二叉树的序列化与反序列化
//输入：root = [1,2,3,null,null,4,5]
//输出：[1,2,3,null,null,4,5]
type Codec struct {
	queue []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string { //DFS前序遍历
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.queue = strings.Split(data, ",")
	return this.helper()
}

func (this *Codec) helper() *TreeNode {
	valStr := this.queue[0]
	this.queue = this.queue[1:]
	if valStr == "#" {
		return nil
	}
	val, _ := strconv.Atoi(valStr)
	return &TreeNode{
		Val:   val,
		Left:  this.helper(),
		Right: this.helper(),
	}
}
