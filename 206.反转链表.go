package main

func main() {

}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		head.Next, head, pre = pre, head.Next, head
	}
	return pre
}
