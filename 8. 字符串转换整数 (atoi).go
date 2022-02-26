package main

import "strconv"

func main() {
	
}
func myAtoi(str string) int {
	begin := true
	ans := []byte{'0'}
	fuhao :=1
	for _, s := range str {
		if s == ' '{
			if begin{
				continue
			}else{
				break
			}
		}else if s== '+'{
			if begin{
				begin = false
			}else{
				break
			}
		}else if s == '-'{
			if begin{
				begin = false
				fuhao = -1
			}else{
				break
			}
		}else if s >='0' && s<='9'{
			ans = append(ans, byte(s))
			begin = false
		}else{
			break
		}
	}
	i,_ := strconv.Atoi(string(ans))
	if i > 2147483647{
		if fuhao ==1{
			i = 2147483647
		}else{
			i = 2147483648
		}
	}
	return i*fuhao
}