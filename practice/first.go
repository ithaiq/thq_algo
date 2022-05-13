package main

import (
	"sort"
)

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
