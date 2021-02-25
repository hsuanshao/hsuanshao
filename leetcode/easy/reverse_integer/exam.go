package main

import "strconv"

//Runtime: 0 ms
//Memory Usage: 2.2 MB

func main() {}

func reverse(x int) int {
	negNum := false
	if x < 0 {
		negNum = true
	}
	str := strconv.Itoa(x)
	revert := ""
	for i := len(str) - 1; i >= 0; i-- {
		if negNum && i == 0 {
			break
		}
		revert += string(str[i])
	}

	res, _ := strconv.Atoi(revert)
	boundary := power(2, 31)
	if res > (boundary - 1) {
		return 0
	}

	if negNum {
		res = 0 - res
		if res < (0 - boundary) {
			return 0
		}
	}

	return res
}

func power(x, t int) int {
	res := 1
	for i := 0; i < t; i++ {
		res = res * x
	}
	return res
}
