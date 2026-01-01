package dsa_01_01_25
//https://leetcode.com/problems/two-sum/

import "sort"

// soln 1 o(n2)
// func twoSum(nums []int, target int) []int {
//     // using a brute force approach

//     for i := 0; i < len(nums); i++ {

//         for j := i+1; j< len(nums); j++ {

//             if nums[i] + nums[j] == target{
//                 return []int{i, j}
//             }
//         }
//     }


//     return []int{}
// }

// soln 2 o(n)
// func twoSum(nums []int, target int) []int {
//     // using a map

//     item := make(map[int]int)

//     for i := 0; i < len(nums); i++ {
//         rem := target - nums[i]

        
//         if val, ok := item[rem]; ok{
//             return []int{val, i}
//         }


//         item[nums[i]] = i
//     }


//     return []int{}
// }

// soln 3 O(logn)
func twoSum(nums []int, target int) []int {

    arr := make([][]int,0, len(nums))

    for i:=0; i < len(nums); i++{
       arr = append(arr, []int{nums[i], i}) 
    }

    sort.Slice(arr, func(i, j int) bool {
        return arr[i][0] < arr[j][0]
    })

    left, right := 0,len(arr)-1

    for left < right {
        sum := arr[left][0] + arr[right][0]

        if sum == target {
            return []int{arr[left][1], arr[right][1]}
        }else if sum < target {
            left++
        }else {
            right--
        }
    }

    return []int{}
}

func Main() {
}