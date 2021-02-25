package main

import "fmt"

/*
Runtime: 4 ms
Memory Usage: 3.2 MB
*/

func main() {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if sum := nums[i] + nums[j]; sum != target {
				continue
			}
			return []int{i, j}
		}
	}
	return []int{}
}
