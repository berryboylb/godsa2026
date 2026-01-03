package dsa_03_01_2026

// https://leetcode.com/problems/contains-duplicate/
import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
	i := 1

	for i < len(nums) {
		if nums[i-1] == nums[i] {
			return true
		}
		fmt.Println("[values]",nums[i-1], nums[i] )

		i++
	}

	return false
}

func containsDuplicate2(nums []int) bool {
	seen := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
        // early return
		if seen[nums[i]] {
			return true
		}

		seen[nums[i]] = true
	}

	return false
}