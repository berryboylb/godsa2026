package dsa_02_01_2026

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

import "fmt"

func maxProfit(prices []int) int {
    profit := 0

	for i := 0; i < len(prices); i++ {

		for j := i + 1; j < len(prices); j++ {
			curr := prices[j] - prices[i]

			if curr > profit {
				profit = curr
			}
		}

	}
	fmt.Println(">>> [profit] <<<", profit)
	return profit
}

func solution2(prices []int) int {
    profit, min := 0, prices[0]

	for i := 1; i < len(prices); i++ {

        curr := prices[i] - min

        if curr > profit {
			profit = curr
        }

        if prices[i] < min {
            min = prices[i]
        }

	}
	fmt.Println(">>> [profit] <<<", profit)
	return profit
}