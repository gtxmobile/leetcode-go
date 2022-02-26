package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	cur := head
	last := &ListNode{-1, head}
	dummy := last
	var pre *ListNode
	for cur != nil {
		tmp := k
		tmp_head := cur
		pre = nil
		for cur != nil && tmp > 0 {
			cur, cur.Next, pre = cur.Next, pre, cur
			tmp--
		}

		if tmp > 0 {
			cur = pre
			pre = nil
			for cur != nil {
				cur, cur.Next, pre = cur.Next, pre, cur
			}

		}
		last.Next = pre
		last = tmp_head
		if tmp >0{
			break
		}
	}
	return dummy.Next
}
