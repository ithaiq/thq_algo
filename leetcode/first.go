package main

import (
	"sort"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.cn/problems/two-sum/
//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
func twoSum(nums []int, target int) []int {
	mapNum := make(map[int]int)
	for k, v := range nums {
		if index, ok := mapNum[target-v]; ok {
			return []int{k, index}
		}
		mapNum[v] = k
	}
	return nil
}

//https://leetcode.cn/problems/add-two-numbers/
//给你两个 非空 的链表，表示两个非负的整数。请你将两个数相加，并以相同形式返回一个表示和的链表。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	tmp := 0
	for l1 != nil || l2 != nil {
		sum := tmp
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		head.Next = &ListNode{Val: sum % 10}
		tmp = sum / 10
		head = head.Next
	}
	if tmp != 0 {
		head.Next = &ListNode{Val: tmp}
	}
	return dummy.Next
}

//https://leetcode.cn/problems/longest-substring-without-repeating-characters/
//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
func lengthOfLongestSubstring(s string) int { ///
	max := 0
	num := make([]byte, 0)
	numMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := numMap[s[i]]; ok {
			for {
				tmp := num[0]
				if len(num) == 1 {
					num = []byte{}
				} else {
					num = num[1:]
				}
				delete(numMap, tmp)
				if tmp == s[i] {
					break
				}
			}

		}
		num = append(num, s[i])
		numMap[s[i]] = i
		if len(num) > max {
			max = len(num)
		}
	}
	return max
}

func lengthOfLongestSubstring2(s string) int { ///
	if len(s) < 2 {
		return len(s)
	}
	left, right, max := 0, 1, 1
	for right < len(s) {
		for i := left; i < right; i++ {
			if s[i] == s[right] {
				left = i + 1
				break
			}
		}
		if right-left+1 > max {
			max = right - left + 1
		}
		right++
	}
	return max
}

//https://leetcode.cn/problems/longest-palindromic-substring/
//给你一个字符串 s，找到 s 中最长的回文子串。
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
func longestPalindrome(s string) string { ///
	if len(s) < 2 {
		return s
	}
	index, maxLen := 0, 0

	getLongestPalindrome := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLen {
				maxLen = right - left + 1
				index = left
			}
			left--
			right++
		}
	}
	for i := 0; i < len(s); i++ {
		getLongestPalindrome(i-1, i+1)
		getLongestPalindrome(i, i+1)
	}
	if maxLen == 0 {
		return s[index : index+1] ///
	}
	return s[index : index+maxLen]
}

//https://leetcode.cn/problems/3sum/
//三数之和
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
func threeSum(nums []int) [][]int { ///
	var result [][]int
	sort.Ints(nums)
	for k := 0; k < len(nums)-2; k++ {
		if k == 0 || nums[k] != nums[k-1] { ///
			i, j := k+1, len(nums)-1
			for i < j {
				sum := nums[i] + nums[j] + nums[k]
				if sum == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
					i++
					j--
					for i < j && nums[i] == nums[i-1] { ///
						i++
					}
					for i < j && nums[j] == nums[j+1] {
						j--
					}
				} else if sum < 0 {
					i++
				} else if sum > 0 {
					j--
				}
			}
		}
	}
	return result
}

//https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
// 删除链表的倒数第 N 个结点
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	head = dummy

	next := head
	for i := 0; i <= n; i++ { ///=
		next = next.Next
	}
	for next != nil {
		head = head.Next
		next = next.Next
	}
	head.Next = head.Next.Next
	return dummy.Next
}

//https://leetcode.cn/problems/valid-parentheses/
//有效的括号
//输入：s = "()[]{}"
//输出：true
func isValid(s string) bool {
	var charMap = map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {

		if v, ok := charMap[s[i]]; ok {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 { ///
				return false
			}
			char := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if char != s[i] {
				return false
			}
		}
	}
	return len(stack) == 0
}

//https://leetcode.cn/problems/merge-two-sorted-lists/
//合并两个有序链表
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			head.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		head = head.Next
	}
	if list1 != nil {
		head.Next = list1
	}
	if list2 != nil {
		head.Next = list2
	}
	return dummy.Next
}

//https://leetcode.cn/problems/reverse-linked-list/
// 反转链表
// p 1 2 3
//   1 nil
//   p head
func reverseList(head *ListNode) *ListNode { ///
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre

		pre = head
		head = next
	}
	return pre
}

//https://leetcode.cn/problems/swap-nodes-in-pairs/
//两两交换链表中的节点
//输入：[1,2,3,4]
//输出：[2,1,4,3]
//c 1 2 3 4
//c 2 1 3 4
//	  c
func swapPairs(head *ListNode) *ListNode { ///
	dummy := &ListNode{}
	dummy.Next = head
	newHead := dummy
	for dummy != nil && dummy.Next != nil && dummy.Next.Next != nil {
		n1 := dummy.Next
		n2 := dummy.Next.Next
		next := n2.Next

		n2.Next = n1
		n1.Next = next
		dummy.Next = n2
		dummy = n1
	}
	return newHead.Next
}

//https://leetcode.cn/problems/group-anagrams/
//字母异位词分组
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
func groupAnagrams(strs []string) [][]string {
	strMap := make(map[[26]int][]string)

	for _, str := range strs {
		arr := [26]int{}
		for i := 0; i < len(str); i++ {
			arr[str[i]-'a'] += 1
		}
		val, ok := strMap[arr]
		if ok {
			val = append(val, str)
		} else {
			val = []string{str}
		}
		strMap[arr] = val ///
	}
	ret := make([][]string, 0, len(strMap))
	for _, v := range strMap {
		ret = append(ret, v)
	}
	return ret
}

//https://leetcode.cn/problems/maximum-subarray/
//最大子数组和
//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
func maxSubArray(nums []int) int {
	maxFun := func(a, b int) int {
		if a < b {
			return b
		} else {
			return a
		}
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = maxFun(dp[i-1]+nums[i], nums[i])
		max = maxFun(max, dp[i])
	}
	return max
}

//https://leetcode.cn/problems/spiral-matrix/
// 螺旋矩阵
//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
func spiralOrder(matrix [][]int) []int { ///
	if len(matrix) == 0 {
		return nil
	}
	left, right, top, down := 0, len(matrix[0])-1, 0, len(matrix)-1
	dir := "right"
	ret := make([]int, 0, len(matrix[0])*len(matrix))
	for left <= right && top <= down {
		if dir == "right" {
			for i := left; i <= right; i++ {
				ret = append(ret, matrix[top][i])
			}
			dir = "down"
			top++
		} else if dir == "down" {
			for i := top; i <= down; i++ {
				ret = append(ret, matrix[i][right])
			}
			dir = "left"
			right--
		} else if dir == "left" {
			for i := right; i >= left; i-- {
				ret = append(ret, matrix[down][i])
			}
			dir = "top"
			down--
		} else if dir == "top" {
			for i := down; i >= top; i-- {
				ret = append(ret, matrix[i][left])
			}
			dir = "right"
			left++
		}
	}
	return ret
}

//https://leetcode.cn/problems/jump-game/
//跳跃游戏
//输入：nums = [2,3,2,1,4]
//输出：true
//解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
func canJump1(nums []int) bool {
	dp := make([]int, len(nums))
	dp[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		maxLen := nums[i] + i
		for j := i + 1; j <= maxLen && j < len(nums); j++ {
			if dp[j] == 1 {
				dp[i] = 1
				break
			}
		}
	}
	return dp[0] == 1
}

func canJump2(nums []int) bool { ///
	maxLen := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= maxLen {
			maxLen = i
		}
	}
	return maxLen == 0
}

//https://leetcode.cn/problems/merge-intervals/
//合并区间
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	}) ///
	cur := intervals[0]
	ret := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= cur[1] {
			if intervals[i][1] >= cur[1] {
				cur[1] = intervals[i][1]
			}
		} else {
			ret = append(ret, cur)
			cur = intervals[i]
		}
	}
	if cur != nil {
		ret = append(ret, cur)
	}
	return ret
}

//https://leetcode.cn/problems/unique-paths/
//不同路径
//一个机器人位于一个 m x n网格的左上角 （起始点在下图中标记为 “Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
//问总共有多少条不同的路径？
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

//https://leetcode.cn/problems/plus-one/
//加一
//输入：digits = [1,2,3]
//输出：[1,2,4]
//解释：输入数组表示数字 123
func plusOne(digits []int) []int {
	ret := make([]int, len(digits))
	v := 1
	for i := len(digits) - 1; i >= 0; i-- {
		ret[i] = (digits[i] + v) % 10
		v = (digits[i] + v) / 10
	}
	if v != 0 {
		fin := []int{v}
		fin = append(fin, ret...)
		return fin
	}
	return ret
}

//https://leetcode.cn/problems/climbing-stairs/
//爬楼梯
//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[n] = dp[n-1] + dp[n-2]
	}
	return dp[n]
}

//https://leetcode.cn/problems/set-matrix-zeroes/
//矩阵置零
//输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
//输出：[[1,0,1],[0,0,0],[1,0,1]]
func setZeroes(matrix [][]int) {

}

//https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
//删除排序链表中的重复元素
//输入：head = [1,1,2,3,3]
//输出：[1,2,3]
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return cur
}

//https://leetcode.cn/problems/reverse-linked-list-ii/
//反转链表 II
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
//[1,2,3,4,5]
//[1,4,3,2,5]
func reverseBetween(head *ListNode, left int, right int) *ListNode { ///
	var pre *ListNode
	cur := head
	for i := 1; i < left; i++ {
		pre = cur
		cur = cur.Next
	}
	pre2, cur2 := pre, cur
	for i := left; i <= right; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	if pre2 != nil {
		pre2.Next = pre
	} else {
		head = pre
	}
	cur2.Next = cur
	return head
}

//https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
//买卖股票的最佳时机
//输入：[7,1,5,3,6,4]
//输出：5
//解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	max := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		val := prices[i] - min
		if val > max {
			max = val
		}
	}
	return max
}

//https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
//买卖股票的最佳时机 II
//给你一个整数数组 prices ，其中prices[i] 表示某支股票第 i 天的价格。
//
//在每一天，你可以决定是否购买和/或出售股票。你在任何时候最多只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
//
//返回 你能获得的 最大 利润。
func maxProfit2(prices []int) int { ///
	if len(prices) < 2 {
		return 0
	}
	val := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			val += prices[i+1] - prices[i]
		}
	}
	return val
}
func maxProfit2_2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	min, max, val := prices[0], prices[0], 0

	for i := 0; i < len(prices)-1; i++ {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		min = prices[i]
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		max = prices[i]
		val += max - min
	}
	return val
}

//https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
// 买卖股票的最佳时机 III
func maxProfit3(prices []int) int {
	dp := make([][]int, 3)
	dp[0] = make([]int, len(prices))
	dp[1] = make([]int, len(prices))
	dp[2] = make([]int, len(prices))

	for i := 1; i < 3; i++ {
		maxProfit := -prices[0]
		for j := 1; j < len(prices); j++ {
			dp[i][j] = max(dp[i][j-1], prices[j]+maxProfit)
			maxProfit = max(maxProfit, dp[i-1][j]-prices[j])
		}
	}
	return dp[2][len(prices)-1]
}

//https://leetcode.cn/problems/gas-station/
//加油站
func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost := 0, 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}
	if totalGas < totalCost {
		return -1
	}
	currentGas, start := 0, 0
	for i := 0; i < len(gas); i++ {
		currentGas = currentGas - cost[i] + gas[i]
		if currentGas < 0 {
			currentGas = 0
			start = i + 1
		}
	}
	return start
}

//https://leetcode.cn/problems/linked-list-cycle/
//环形链表
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil { ///
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

//https://leetcode.cn/problems/linked-list-cycle-ii/
//环形链表 II
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast, isCycle := head, head, false
	for fast.Next != nil && fast.Next.Next != nil { ///
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			isCycle = true
			break
		}
	}
	if !isCycle {
		return nil
	}
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}

//https://leetcode.cn/problems/maximum-product-subarray/
//乘积最大子数组
//输入: nums = [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
func maxProduct(nums []int) int {
	maxDp := make([]int, len(nums))
	minDp := make([]int, len(nums))
	maxDp[0] = nums[0]
	minDp[0] = nums[0]
	maxV := nums[0]
	_max := func(a, b, c int) (v int) {
		if a > b {
			v = a
		} else {
			v = b
		}
		if v < c {
			v = c
		}
		return
	}
	_min := func(a, b, c int) (v int) {
		if a > b {
			v = b
		} else {
			v = a
		}
		if v > c {
			v = c
		}
		return
	}
	for i := 1; i < len(nums); i++ {
		maxDp[i] = _max(nums[i], nums[i]*maxDp[i-1], nums[i]*minDp[i-1])
		minDp[i] = _min(nums[i], nums[i]*maxDp[i-1], nums[i]*minDp[i-1])
		maxV = max(maxV, maxDp[i])
	}
	return maxV
}

//https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/
//寻找旋转排序数组中的最小值
//你必须设计一个时间复杂度为 O(log n) 的算法解决此问题
//输入：nums = [3,4,5,1,2]
//输出：1
//解释：原数组为 [1,2,3,4,5] ，旋转 3 次得到输入数组。
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := 0, len(nums)-1
	if nums[right] > nums[0] {
		return nums[0]
	}
	for left < right {
		mid := (right-left)/2 + left
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		if nums[mid] > nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nums[0]
}

//https://leetcode.cn/problems/intersection-of-two-linked-lists/
//相交链表
//给你两个单链表的头节点 headA 和 headB
//请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	n1 := headA
	n2 := headB
	for n1 != n2 { //包含不相交的情形
		if n1 == nil {
			n1 = headB
		} else {
			n1 = n1.Next
		}
		if n2 == nil {
			n2 = headA
		} else {
			n2 = n2.Next
		}
	}
	return n1
}

//https://leetcode.cn/problems/repeated-dna-sequences/
//重复的DNA序列
//返回所有在 DNA 分子中出现不止一次的 长度为 10 的序列(子字符串)。你可以按 任意顺序 返回答案。
//输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
//输出：["AAAAACCCCC","CCCCCAAAAA"]

func findRepeatedDnaSequences(s string) []string {
	return nil
}

//https://leetcode.cn/problems/house-robber/
//打家劫舍
//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//输入：[1,2,3,1]
//输出：4
//解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//    偷窃到的最高金额 = 1 + 3 = 4 。

func rob(nums []int) int {

}

//https://leetcode.cn/problems/sort-array-by-parity/
//按奇偶排序数组
//给你一个整数数组 nums，将 nums 中的的所有偶数元素移动到数组的前面，后跟所有奇数元素。
//输入：nums = [3,1,2,4]
//输出：[2,4,3,1]
func sortArrayByParity(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	o, j := 0, len(nums)-1
	for o < j {
		oV := nums[o] % 2
		jv := nums[j] % 2
		if oV == 0 && jv == 0 {
			o++
		} else if oV == 0 && jv == 1 {
			o++
			j--
		} else if oV == 1 && jv == 0 {
			nums[o], nums[j] = nums[j], nums[o]
			o++
			j--
		} else if oV == 1 && jv == 1 {
			j--
		}
	}
	return nums
}

//https://leetcode.cn/problems/sort-array-by-parity-ii/
//按奇偶排序数组 II
//给定一个非负整数数组nums，nums 中一半整数是 奇数 ，一半整数是 偶数 。
//
//对数组进行排序，以便当nums[i] 为奇数时，i也是 奇数 ；当nums[i]为偶数时， i 也是 偶数 。
//输入：nums = [4,2,5,7]
//输出：[4,5,2,7]
func sortArrayByParityII(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	o, j := 0, 1
	for ; o < len(nums); o += 2 {
		if nums[o]%2 == 1 {
			for j < len(nums) && nums[j]%2 == 1 {
				j += 2
			}
			nums[o], nums[j] = nums[j], nums[o]
		}
	}
	return nums
}

//https://leetcode.cn/problems/odd-even-linked-list/
//奇偶链表
//输入: head = [2,1,3,5,6,4,7]
//输出: [2,3,6,7,1,5,4]
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	j := head
	o := head.Next
	oH := head.Next
	for o != nil && o.Next != nil {
		j.Next = j.Next.Next
		j = j.Next
		o.Next = o.Next.Next
		o = o.Next
	}
	j.Next = oH
	return head
}

//https://leetcode.cn/problems/fruit-into-baskets/
//水果成篮
//相连的果树只能采摘2棵
//输入：fruits = [1,2,3,2,2]
//输出：4
//解释：可以采摘 [2,3,2,2] 这四棵树。
//如果从第一棵树开始采摘，则只能采摘 [1,2] 这两棵树。
func totalFruit(fruits []int) int {
	if len(fruits) <= 2 {
		return len(fruits)
	}
	fM := make(map[int]int)
	max := 0
	j := 0
	for i := 0; i < len(fruits); i++ {
		fM[fruits[i]] = i
		if len(fM) > 2 {
			minIndex := len(fruits)
			minK := len(fruits)
			for k, v := range fM {
				if v < minIndex {
					minIndex = v
					minK = k
				}
			}
			delete(fM, minK)
			j = minIndex + 1
		}
		if i-j+1 > max { //每次更新
			max = i - j + 1
		}
	}
	return max
}

//https://leetcode.cn/problems/middle-of-the-linked-list/
//链表的中间结点
func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//https://leetcode.cn/problems/backspace-string-compare/
//比较含退格的字符串
//输入：s = "ab##", t = "c#d#"
//输出：true
//解释：s 和 t 都会变成 ""。
func backspaceCompare(s string, t string) bool {
	i := len(s) - 1
	j := len(t) - 1
	backI := 0
	backJ := 0
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				backI++
				i--
			} else if backI > 0 {
				backI--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				backJ++
				j--
			} else if backJ > 0 {
				backJ--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 && s[i] != t[j] {
			return false
		} else if i >= 0 && j < 0 {
			return false
		} else if i < 0 && j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}

//https://leetcode.cn/problems/rotate-string/
// 旋转字符串
//输入: s = "abcde", goal = "cdeab"
//输出: true
func rotateString(s string, goal string) bool {
	newStr := s + s
	for i := 0; i < len(s); i++ {
		if goal == newStr[i:len(s)+i] {
			return true
		}
	}
	return false
}

//https://leetcode.cn/problems/number-of-islands/
// 岛屿数量
//给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
//此外，你可以假设该网格的四条边均被水包围。
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
func numIslands(grid [][]byte) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				max++
				numIslands_dfs(grid, i, j)
			}
		}
	}
	return max
}
func numIslands_dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	dfs2(grid, i-1, j)
	dfs2(grid, i+1, j)
	dfs2(grid, i, j-1)
	dfs2(grid, i, j+1)
	return
}

//https://leetcode.cn/problems/max-area-of-island/
//岛屿的最大面积
//给你一个大小为 m x n 的二进制矩阵 grid 。
//岛屿是由一些相邻的1(代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设grid 的四个边缘都被 0（代表水）包围着。
//岛屿的面积是岛上值为 1 的单元格的数目。
//计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。

//输入：grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
//输出：6
//解释：答案不应该是 11 ，因为岛屿只能包含水平或垂直这四个方向上的 1 。
func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				count := dfs(grid, i, j)
				if max < count {
					max = count
				}
			}
		}
	}
	return max
}
func dfs(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	count := 1
	count += dfs(grid, i-1, j)
	count += dfs(grid, i+1, j)
	count += dfs(grid, i, j-1)
	count += dfs(grid, i, j+1)
	return count
}

//https://leetcode.cn/problems/battleships-in-a-board/
//甲板上的战舰
//输入：board = [["X",".",".","X"],[".",".",".","X"],[".",".",".","X"]]
//输出：2
func countBattleships(board [][]byte) int {
	max := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' {
				max++
				dfs2(board, i, j)
			}
		}
	}
	return max
}
func dfs2(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 'X' {
		return
	}
	grid[i][j] = 0
	dfs2(grid, i-1, j)
	dfs2(grid, i+1, j)
	dfs2(grid, i, j-1)
	dfs2(grid, i, j+1)
	return
}

//https://leetcode.cn/problems/valid-palindrome-ii/
//验证回文字符串 Ⅱ
//给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
//输入: s = "aba"
//输出: true
func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}
		left++
		right--
	}
	return true
}
func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

//https://leetcode.cn/problems/add-two-numbers-ii/
//链表两数相加 II
//输入：l1 = [7,2,4,3], l2 = [5,6,4]
//输出：[7,8,0,7]
func addTwoNumbersx(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1, stack2 []int
	var cur *ListNode
	carry := 0
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	for len(stack1) != 0 || len(stack2) != 0 {
		sum := carry
		if len(stack1) != 0 {
			v := stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
			sum += v
		}
		if len(stack2) != 0 {
			v := stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
			sum += v
		}
		node := &ListNode{Val: sum % 10}
		node.Next = cur
		carry = sum / 10
		cur = node
	}
	if carry > 0 {
		node := &ListNode{Val: carry}
		node.Next = cur
		cur = node
	}
	return cur
}

//https://leetcode.cn/problems/move-zeroes/
//移动零
//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]
func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}

//https://leetcode.cn/problems/product-of-array-except-self/
//除自身以外数组的乘积
//输入: nums = [1,2,3,4]
//输出: [24,12,8,6]
//1 2 3 4 5
//1 1 2 6 24
//120  60  40  30  24
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(result); i++ {
		result[i] = 1
	}
	pre := 1
	for i := 0; i < len(nums); i++ {
		result[i] = result[i] * pre
		pre = pre * nums[i]
	}
	pre = 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * pre
		pre = pre * nums[i]
	}
	return result
}
