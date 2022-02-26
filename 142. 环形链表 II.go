package main

func main() {

}
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	one := head
	two := head
	n3 := head
	for two != nil && two.Next != nil {
		one = one.Next
		two = two.Next.Next
		if one == two {
			for n3 != one {
				n3 = n3.Next
				one = one.Next
			}
			return one
		}
	}
	return nil
}
