package main

import "fmt"

// Runtime: 8 ms
// Memory Usage: 3.1 MB

func main() {
	res := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	res := 0

	for l, r := 0, 0; r < len(s); r++ {
		if idx, ok := m[s[r]]; ok {
			if l < idx {
				l = idx
			}
		}

		if res < (r - l + 1) {
			res = r - l + 1
		}
		m[s[r]] = r + 1
	}
	return res
}
