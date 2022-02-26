package main

func main() {
	
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil{
		return true
	}
	l := 0
	cur := head
	for cur != nil{
		cur = cur.Next
		l++
	}
	mid := l/2
	cur = head
	var pre *ListNode
	for mid >0{
		cur.Next,cur,pre = pre, cur.Next,cur
		mid--
	}
	left := pre
	right := cur
	if l%2 == 1{
		right = cur.Next
	}
	mid = l/2
	for mid > 0{
		if left.Val != right.Val{
			return false
		}
		left, right = left.Next,right.Next
		mid--
	}
	return true
}