// DSA Daily — 2026-01-06
// Problem: https://leetcode.com/problems/maximum-subarray/description/

package dsa06012026

func solution(nums []int) int {
    curr, max := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if curr < 0 {
			curr = 0
		}
		curr += nums[i]
		if curr > max {
			max = curr
		}
	}

	return max
}

func main() {}
