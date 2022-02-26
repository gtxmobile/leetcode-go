package main

func main() {
	println(countPrimes1(10000))
}
func countPrimes1(n int) int {
	if n<=2{
		return 0
	}
	notprime := make([]bool, n)
	for i := 2;i*i<n;i++{
		if notprime[i]{
			continue
		}
		for j:= i*i;j<n;j+=i{
			if !notprime[j]{
				notprime[j] = true
			}
		}
	}
	cnt :=0
	for _,v :=range notprime[2:]{
		if !v{
			cnt++
		}
	}
	return cnt
}
func countPrimes(n int) int {
	if n<=2{
		return 0
	}
	notprime := make([]bool, n)
	cnt := n/2
	for i := 3;i*i<n;i+=2{
		if notprime[i]{
			continue
		}
		for j:= i*i;j<n;j+=2*i{
			if !notprime[j]{
				notprime[j] = true
				cnt--
			}
		}
	}
	return cnt
}
