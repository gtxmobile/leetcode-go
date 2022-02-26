package main

import "strings"

func main() {

}
func reverseWords(s string) string {
	var ret []string
	raw := strings.Split(s, " ")
	l := len(raw)
	for i := l - 1; i >= 0; i-- {
		if len(raw[i]) > 0 {
			ret = append(ret, raw[i])
		}
	}
	return strings.Join(ret, " ")
}
