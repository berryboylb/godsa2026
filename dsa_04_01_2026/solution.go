// DSA Daily — 2026-01-04
// Problem: https://leetcode.com/problems/contains-duplicate-ii/

package dsa04012026
import (
	"math"
)

func solution(nums []int, k int) bool {
	 // my name is tenshi and i'm scared to fail so i just give up

  // find duplicates
  // compare indexes
  // get absolute
  // compare with k

  for i:=0; i <len(nums); i++ {
    for j := i+1; j < len(nums); j++{
        if nums[i] == nums[j]{
            if int(math.Abs(float64(i-j))) <= k{
                return true
            }
        }
    }
  }

  return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
  // we need to use one loop so we need another data structure

  seen := make(map[int]int)

  for i := 0; i < len(nums); i++ {
    if val, ok := seen[nums[i]]; ok {
        if int(math.Abs(float64(i-val))) <= k {
			return true
		}
    }
    seen[nums[i]] = i
  }

  return false
}

func main() {}
