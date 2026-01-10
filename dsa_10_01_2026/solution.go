// DSA Daily — 2026-01-10
// Problem: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

package dsa10012026

import "fmt"

func solution(numbers []int, target int) []int {
    left, right := 0, len(numbers)-1

	for left <= right {
		sum := numbers[left] + numbers[right]
		if sum > target {
			right -= 1
		} else if sum < target {
			left += 1
		} else {
			fmt.Println("[value]", []int{left + 1, right + 1})
			return []int{left + 1, right + 1}
		}
	}

	return []int{}
}

func main() {}
