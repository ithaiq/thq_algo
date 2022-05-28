package algo

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.cn/problems/reverse-string/
//反转字符串
//输入：s = ["h","e","l","l","o"]
//输出：["o","l","l","e","h"]
func reverseString(s []byte) {
	res := make([]byte, 0)
	reverse(s, 0, &res)
	for i := 0; i < len(s); i++ {
		s[i] = res[i]
	}
}
func reverse(s []byte, i int, res *[]byte) {
	if i == len(s) {
		return
	}
	reverse(s, i+1, res)
	*res = append(*res, s[i])
}

//https://leetcode.cn/problems/swap-nodes-in-pairs/
//两两交换链表中的节点
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//[2,1,4,3]
//1        2        3         4
//head    next    nextHead

func swapPairs(head *ListNode) *ListNode {
	return helper(head)
}
func helper(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nextHead := head.Next.Next
	next := head.Next
	next.Next = head
	head.Next = helper(nextHead)
	return next
}

//递归反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
