package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	bundles := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&bundles[i])
	}
	fmt.Println(getMaxAmount(bundles))
}

// TIME: O(n)
// SPACE: O(n)
// Standard Solution
func getMaxAmount(bundles []int) int {
	maxAmount := 0
	for _, amount := range bundles {
		maxAmount = max(maxAmount, amount)
	}

	sum := make([]int, maxAmount+1)
	for _, amount := range bundles {
		sum[amount] += amount
	}

	dp := make([]int, maxAmount+1)
	dp[0] = 0
	dp[1] = sum[1]

	for i := 2; i <= maxAmount; i++ {
		dp[i] = max(sum[i]+dp[i-2], dp[i-1])
	}
	return dp[maxAmount]
}

//func getMaxAmount(bundles []int) int {
//	countedNums := make(map[int]int)
//	var distinctNums []int
//	for _, num := range bundles {
//		value, ok := countedNums[num]
//		if !ok {
//			countedNums[num] = num
//			distinctNums = append(distinctNums, num)
//		} else {
//			countedNums[num] = value + num
//		}
//	}
//	sort.Ints(distinctNums)
//	value, _ := countedNums[distinctNums[0]]
//
//	prev, curr := 0, value
//	for i := 1; i < len(distinctNums); i++ {
//		value, _ = countedNums[distinctNums[i]]
//		// if we need to consider deletion
//		if distinctNums[i-1]+1 == distinctNums[i] {
//			prev, curr = max(prev, curr), max(prev+value, curr)
//		} else {
//			prev, curr = curr, curr+value
//		}
//	}
//
//	return max(prev, curr)
//}
