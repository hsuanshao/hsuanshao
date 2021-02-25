package main

import "fmt"

/*
Runtime: 0 ms
Memory Usage: 2.4 MB
*/

func main() {
	longStr := longestCommonPrefix([]string{"ab", "a"})
	fmt.Println(longStr)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	// get shortest string length
	var shortLen *int
	t := len(strs[0])
	shortLen = &t
	for i := 1; i < len(strs); i++ {
		tmpLen := len(strs[i])
		if tmpLen < *shortLen {
			shortLen = &tmpLen
		}
	}

	if *shortLen == 0 {
		return ""
	}

	idx := 0

	for i := 1; i <= *shortLen; i++ {
		for s := 1; s < len(strs); s++ {
			if string(strs[0][0:i]) != string(strs[s][0:i]) {
				return string(strs[0][0 : i-1])
			}
			idx = i
		}
	}

	return string(strs[0][0:idx])
}
