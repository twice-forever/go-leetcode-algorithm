package problem

import (
	"strings"
)

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	ansSlice := []string{}
	for _, v := range paths {
		switch v {
		case ".":
			continue
		case "..":
			if len(ansSlice) > 0 {
				ansSlice = ansSlice[:len(ansSlice)-1]
			}
		default:
			if v != "" {
				ansSlice = append(ansSlice, v)
			}
		}
	}
	ans := "/"
	for i, v := range ansSlice {
		if i == len(ansSlice)-1 {
			ans += v
			continue
		}
		ans += v + "/"
	}
	return ans
}
