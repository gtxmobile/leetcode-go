package main

func main() {

}
func longestCommonPrefix(strs []string) string {
	ret := []byte("")
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	minlen := len(strs[0])
	for _, s := range strs {
		curlen := len(s)
		if curlen < minlen {
			minlen = curlen
		}
	}
	for i := 0; i < minlen; i++ {
		a := strs[0][i]
		for _, s := range strs[1:] {
			if a != s[i] {
				return string(ret)
			}
		}
		ret = append(ret, a)
	}
	return string(ret)
}
