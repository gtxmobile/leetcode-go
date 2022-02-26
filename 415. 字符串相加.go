package main

func main() {
	
}
func addStrings(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	if l1 > l2 {
		return addStrings(num2, num1)
	}
	var lin byte = '0'
	le := l1
	ls := l2
	ret := []byte{}
	numyu := num2
	var c uint8=0

	for i:=1;i <= le;i++{
		c = addByte(num1[l1-i], num2[l2-i],c,&ret)
	}
	for j := le + 1; j <= ls; j++ {
		c = addByte(numyu[ls-j], '0', c,&ret)
	}
	if c > 0 {
		ret = append(ret, c+lin)
	}
	lend := len(ret)
	ans := []byte{}
	for k := 1; k <= lend; k++ {
		ans = append(ans, ret[lend-k])
	}
	return string(ans)
}

func addByte(n1 byte, n2 byte , c uint8, ret *[]byte) uint8 {
	he := ( n1- '0') + ( n2- '0') + c
	if he >= 10{
		c = 1
		he = he -10
	} else{
		c = 0
	}
	*ret = append(*ret, he+'0')
	return c
}
