// DSA Daily — 2026-01-08
// Problem: https://leetcode.com/problems/maximum-product-subarray/description/

package dsa08012026
import "math"



func solution(nums []int) int {
    if len(nums) == 0 {
		return 0
	}
	min, max, res := nums[0], nums[0], nums[0]


	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		temp := int(math.Max(float64(curr), math.Max(float64(max*curr), float64(min*curr))))
		min = int(math.Min(float64(curr), math.Min(float64(min*curr), float64(max*curr))))
		max = temp
		res = int(math.Max(float64(res), float64(max)))
	}

	return res
}

func main() {}
