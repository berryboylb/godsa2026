// DSA Daily — 2026-01-11
// Problem: https://leetcode.com/problems/3sum/

package dsa11012026
import ("fmt"; "sort")

func threeSum(nums []int) [][]int {
    res := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			res = twoSum(nums, i, res)
		}

		
	}

	fmt.Println("[data]", res)

	return res
}

func twoSum(nums []int, i int, result [][]int) [][]int {
	left, right := i+1, len(nums)-1

	for left < right {
		sum := nums[i] + nums[left] + nums[right]
		if sum < 0 {
			left++
		} else if sum > 0 {
			right--
		} else {
			result = append(result, []int{nums[i], nums[left], nums[right]})
			left++
			right--

			// skip duplicates
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		}
	}

	return result
}

func main() {}
