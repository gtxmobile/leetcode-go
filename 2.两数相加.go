package main
import "fmt"
func main() {
	fmt.Println(addTwoNumbers([]int{2, 7, 11, 15},9))
}
type ListNode struct {
	Val int
	Next *ListNode
}
func getVal(l3 *ListNode) int{
	if l3 != nil {
		return l3.Val
	}else{
		return 0
	}
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	head := ret
	c := 0
	for l1 != nil || l2 != nil{
		x := getVal(l1)
		y := getVal(l2)
		he := x + y + c
		c = he /10
		ret.Next = &ListNode{he%10, nil}
		ret = ret.Next
		if l1 != nil{
			l1 = l1.Next
		}
		if l2 != nil{
			l2 = l2.Next
		}
	}
	if c >0 {
		ret.Next = new(ListNode)
		ret.Next.Val = c
	}
	return head.Next
}
