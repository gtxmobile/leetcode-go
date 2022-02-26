package main

import "fmt"

func main() {
	fmt.Print(letterCombinations("23"))

}
func letterCombinations(digits string) []string {
	d := map[int][]string{
		2:{"a","b","c"},
		3:{"d","e","f"},
		4:{"g","h","i"},
		5:{"j","k","l"},
		6:{"m","n","o"},
		7:{"p","q","r","s"},
		8:{"t","u","v"},
		9:{"w","x","y","z"},
	}
	if digits ==""{
		return []string{}
	}
	ans := []string{""}
	for _,i := range digits{
		var tmp []string
		for _,j := range d[int(i-'0')]{
			for _,p := range ans{
				tmp = append(tmp, p+j)
			}
		}
		ans = tmp
	}
	return ans
}