package main

func main() {
	
}
func countAndSay(n int) string {
	if n == 1{
		return "1"
	}
	s := []byte{'1'}
	for i:=1;i<n;i++{
		pre := s[0]
		var cnt byte=1
		tmp := []byte{}
		for j:=1;j< len(s);j++{
			if s[j] == pre{
				cnt++
			}else{
				tmp = append(tmp, '0'+cnt)
				tmp = append(tmp, pre)
				pre = s[j]
				cnt = 1
			}
		}
		tmp = append(tmp, '0'+cnt)
		tmp = append(tmp, pre)
		s = tmp
	}
	return string(s)
}