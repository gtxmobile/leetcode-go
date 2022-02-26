package main

func main() {
	
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := ListNode{-1,head}
	slow, fast := &dummy, &dummy
	for n>0 {
		fast = fast.Next
		n--
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}