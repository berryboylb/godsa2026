// DSA Daily — 2026-01-05
// Problem: https://leetcode.com/problems/product-of-array-except-self/description/

package dsa05012026

func solution(nums []int) []int {
	// implement here
	res := make([]int, len(nums))
	pre, post := 1, 1

	for i := 0; i < len(nums); i++ {
        res[i] = pre
        pre = pre * nums[i]
	}

    for j :=len(nums)-1; j >=0; j-- {
        res[j] = post * res[j]
        post = post * nums[j]
    }



	return res
}

func main() {}
