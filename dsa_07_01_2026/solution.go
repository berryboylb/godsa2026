// DSA Daily — 2026-01-07
// Problem: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

package dsa07012026

func findMin(nums []int) int {
    high, low := len(nums)-1, 0

	for low < high {
		mid := (low + high) / 2

		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

func main() {}
