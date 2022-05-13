package main

import (
	"strconv"
)

//https://leetcode-cn.com/problems/min-stack/
//设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
type MinStack struct {
	min   []int
	stack []int
}

func (this *MinStack) Push(x int) {
	min := this.GetMin()
	if x < min {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, min)
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 1 << 31
	}
	min := this.min[len(this.min)-1]
	return min
}

//https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
//波兰表达式计算 > 输入: ["2", "1", "+", "3", "*"] > 输出: 9
func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	var stack []int
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return -1
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result := 0
			switch tokens[i] {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			}
			stack = append(stack, result)
		default:
			val, _ := strconv.Atoi(tokens[i])
			stack = append(stack, val)
		}
	}
	return stack[0]
}

//https://leetcode-cn.com/problems/decode-string/
//给定一个经过编码的字符串，返回它解码后的字符串。 s = "3[a]2[bc]", 返回 "aaabcbc". s = "3[a2[c]]", 返回 "accaccacc". s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			temp := make([]byte, 0)
			for len(stack) != 0 && stack[len(stack)-1] != '[' {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				temp = append(temp, v)
			}
			// pop '['
			stack = stack[:len(stack)-1]
			// pop num
			idx := 1
			for len(stack) >= idx && stack[len(stack)-idx] >= '0' && stack[len(stack)-idx] <= '9' {
				idx++
			}
			// 注意索引边界
			num := stack[len(stack)-idx+1:]
			stack = stack[:len(stack)-idx+1]
			count, _ := strconv.Atoi(string(num))
			for j := 0; j < count; j++ {
				// 把字符正向放回到栈里面
				for j := len(temp) - 1; j >= 0; j-- {
					stack = append(stack, temp[j])
				}
			}
		default:
			stack = append(stack, s[i])

		}
	}
	return string(stack)
}

//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
//给定一个二叉树，返回它的中序遍历。
func inorderTraversal3(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, val.Val)
		root = root.Right
	}
	return result
}

//https://leetcode-cn.com/problems/clone-graph/
//给你无向连通图中一个节点的引用，请你返回该图的深拷贝（克隆）。
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return clone(node, visited)
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if v, ok := visited[node]; ok {
		return v
	}
	newNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	visited[node] = newNode
	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = clone(node.Neighbors[i], visited)
	}
	return newNode
}

//https://leetcode-cn.com/problems/number-of-islands/
//给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && numIslands_dfs(grid, i, j) >= 1 {
				count++
			}
		}
	}
	return count
}

func numIslands_dfs(grid [][]byte, i int, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == '1' {
		grid[i][j] = 0
		return numIslands_dfs(grid, i-1, j) +
			numIslands_dfs(grid, i, j-1) +
			numIslands_dfs(grid, i+1, j) +
			numIslands_dfs(grid, i, j+1) + 1
	}
	return 0
}

//https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。 求在该柱状图中，能够勾勒出来的矩形的最大面积
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	stack := make([]int, 0)
	max := 0
	for i := 0; i <= len(heights); i++ {
		var cur int
		if i == len(heights) {
			cur = 0
		} else {
			cur = heights[i]
		}

	}
}
