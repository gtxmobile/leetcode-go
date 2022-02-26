package main

func main() {
	
}
func getPermutation(n int, k int) string {
	fact := make([]byte, n)
	sub := 1
	for i := 1; i < n; i++ {
		sub *= i
		fact[i-1] = byte(i)
	}
	fact[n-1] = byte(n)
	ans := []byte{}
	k--
	for k > 0 && n>0{
		idx := k / sub
		ans = append(ans, '0'+fact[idx])
		k = k % sub
		n--
		sub /= n
		pop(idx, &fact)
	}
	if k == 0{
		for _,v := range fact{
			ans = append(ans, '0'+v)
		}
	}
	println(string(ans))
	return string(ans)
}
func pop(i int, a *[]byte) {
	// Remove the element at index i from a.
	copy((*a)[i:], (*a)[i+1:]) // Copy last element to index i.
	(*a)[len(*a)-1] = 0        // Erase last element (write zero value).
	*a = (*a)[:len(*a)-1]      // Truncate slice.
}