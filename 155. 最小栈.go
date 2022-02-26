package main

func main() {
	
}
type MinStack struct {
	Mstack []int
	Stack  []int
	mintop int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var minstack MinStack
	return minstack
}

func (this *MinStack) Push(x int) {
	if len(this.Stack) == 0 || this.mintop>x{
		this.mintop = x
	}
	this.Mstack = append(this.Mstack, this.mintop)
	this.Stack = append(this.Stack, x)
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.Mstack = this.Mstack[:len(this.Mstack)-1]
	if len(this.Mstack) >0 && this.Mstack[len(this.Mstack)-1] > this.mintop{
		this.mintop = this.Mstack[len(this.Mstack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.mintop
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */