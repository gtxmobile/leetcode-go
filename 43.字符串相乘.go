package main

func main() {
	
}
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 =="0"{
		return "0"
	}
	l1 := len(num1)
	ans := "0"
	for i := l1 - 1; i >= 0; i-- {
		one := mulByte(num1[i], num2,l1-i-1)
		println(one)
	}
	return ans
}
func mulByte(n1 byte, n2 string, cnt0 int) string{
	var ret []byte
	intn1 := n1-'0'
	ln2 := len(n2)
	var c byte=0
	for i:=ln2-1;i>=0;i-- {
		ji := intn1 * (n2[i]-'0')+c
		c = ji /10
		ret = append(ret, ji%10+'0')
	}
	if c > 0 {
		ret = append(ret, c+'0')
	}
	lend := len(ret)
	ans := []byte{}
	for k := 1; k <= lend; k++ {
		ans = append(ans, ret[lend-k])
	}
	for i:=0;i<cnt0;i++{
		ans = append(ans, '0')
	}
	return string(ans)
}