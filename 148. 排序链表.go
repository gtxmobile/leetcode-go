package main

func main() {

}
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, nil}
	one, two := head, head.Next
	for two != nil && two.Next != nil {
		one, two = one.Next, two.Next.Next
	}
	mid := one.Next
	one.Next = nil
	left := sortList(head)
	right := sortList(mid)
	cur := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	} else {
		cur.Next = right
	}
	return dummy.Next
}
