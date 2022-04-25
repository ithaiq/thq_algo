package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

func deleteDuplicates1(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
func deleteDuplicates2(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		for cur.Next != nil && cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}

//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
//给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现的数字。

func deleteDuplicates(head *ListNode) *ListNode {
	var dummy = &ListNode{}
	dummy.Next = head
	head = dummy
	rmVal := 0
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			rmVal = head.Next.Val
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}

//https://leetcode-cn.com/problems/reverse-linked-list/
//反转一个单链表。
//A->B->C
//B->A->nil
//C->B->A->nil
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

//https://leetcode-cn.com/problems/reverse-linked-list-ii/
//反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	var dummy = &ListNode{}
	dummy.Next = head
	head = dummy

	var pre *ListNode
	var i = 0
	for ; i < m; i++ {
		pre = head
		head = head.Next
	}
	var j = i
	var mid = head
	var last *ListNode
	for ; head != nil && j <= n; i++ {
		next := head.Next
		head.Next = last
		last = head
		head = next
	}
	pre.Next = last
	mid.Next = head //
	return dummy.Next
}

//https://leetcode-cn.com/problems/merge-two-sorted-lists/
//将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy = &ListNode{}
	head := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	for l1 != nil {
		head.Next = l1
		l1 = l1.Next
		head = head.Next
	}
	for l2 != nil {
		head.Next = l2
		l2 = l2.Next
		head = head.Next
	}
	return dummy.Next
}

//https://leetcode-cn.com/problems/partition-list/
//给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	var dummy = &ListNode{}
	var tailDummy = &ListNode{}
	dummy.Next = head
	head = dummy
	tail := tailDummy
	for head.Next != nil {
		if head.Next.Val >= x {
			tail.Next = head.Next
			head.Next = head.Next.Next
			tail = tail.Next
		} else {
			head = head.Next //
		}
	}
	tail.Next = nil
	head.Next = tailDummy.Next
	return dummy.Next
}

//https://leetcode-cn.com/problems/sort-list/
//在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMiddle(head)
	tail := mid.Next
	mid.Next = nil
	left := mergeSort(head)
	right := mergeSort(tail)
	return mergeTwoLists(left, right)
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//https://leetcode-cn.com/problems/reorder-list/
//重排链表：给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
//L0 → L1 → … → Ln - 1 → Ln
//请将其重新排列后变为：
//
//L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
func reorderList(head *ListNode) {
	// 思路：找到中点断开，翻转后面部分，然后合并前后两个链表
	if head == nil {
		return
	}
	mid := findMiddle(head)
	tail := reverseList(mid.Next)
	mid.Next = nil
	head = mergeTwoLists(head, tail)
}

//https://leetcode-cn.com/problems/linked-list-cycle/
//给定一个链表，判断链表中是否有环
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

//https://leetcode-cn.com/problems/linked-list-cycle-ii/
//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		if fast == slow {
			fast = head
			slow = slow.Next
			for fast != slow {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}

//https://leetcode-cn.com/problems/palindrome-linked-list/
//请判断一个链表是否为回文链表。
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	new := reverseList(slow.Next)
	slow.Next = nil
	for head != nil && new != nil {
		if head.Val != new.Val {
			return false
		}
		head = head.Next
		new = new.Next
	}
	return true
}
