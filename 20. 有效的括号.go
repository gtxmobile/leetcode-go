package main

func main() {
	
}
func isValid(s string) bool {
	var stack []int32
	for _, b := range s{
		l :=len(stack)
		if  l== 0{
			stack = append(stack, b)
			continue
		}
		if (b == ']' && stack[l-1] == '[') || (b == ')' && stack[l-1] == '(') ||(b == '}' && stack[l-1] == '{') {
			stack = stack[:l-1]
			continue
		}
		stack = append(stack, b)

	}
	return len(stack) == 0
}
