// DSA Daily — 2026-01-17
// Problem: https://leetcode.com/problems/container-with-most-water/description/

package dsa17012026
import ("fmt")

func maxArea(height []int) int {
    max, left, right := 0, 0, len(height)-1

	for left <= right {
		currHeight := 0
		moveRight := false
		if height[left] > height[right] {
			currHeight = height[right]
		} else {
			moveRight = true
			currHeight = height[left]
		}
		width := right - left
		area := currHeight * width

		if area > max {
			max = area
		}

		if moveRight {
			left++
		} else {
			right--
		}
	}

	fmt.Println("[max]", max)

	return max
}

func main() {}
