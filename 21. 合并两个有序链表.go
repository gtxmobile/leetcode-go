package main

func main() {
	
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := ListNode{-1,nil}
	last := &dummy
	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val{
			last.Next = l1
			last = last.Next
			l1 = l1.Next
		} else{
			last.Next = l2
			last = last.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		last.Next =  l1
	}
	if l2 != nil {
		last.Next = l2
	}
	return dummy.Next
}