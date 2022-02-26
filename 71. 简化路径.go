package main

import "strings"

func main() {

}
func simplifyPath(path string) string {
	var ans []string
	for _, v := range strings.Split(path, "/") {
		if v == "." || v == "" {
			continue
		} else if v == ".." {
			if len(ans) > 0 {
				ans = ans[:len(ans)-1]
			}
		} else {
			ans = append(ans, v)
		}
	}
	return "/" + strings.Join(ans, "/")
}
