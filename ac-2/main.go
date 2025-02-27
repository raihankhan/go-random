package main

import (
	"fmt"
	"sort"
)

func main() {
	startTime := []int{1, 5, 8}
	endTime := []int{3, 9, 12}
	profit := []int{20, 40, 30}

	fmt.Println(jobScheduling(startTime, endTime, profit))
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	// initiate a dp array where index is start time, value is the maxProfit
	dp := make([]int, len(startTime)+1)
	// combine the arrays in a common job array, sort the array based on start time
	job := make([][3]int, len(startTime))
	for i := 0; i < len(startTime); i++ {
		job[i][0], job[i][1], job[i][2] = startTime[i], endTime[i], profit[i]
	}
	sort.Slice(job, func(i, j int) bool { return job[i][0] < job[j][0] })

	// sort the startTime array too, now both job[][0] array and sort array are similarly indexed
	sort.Ints(startTime)

	for i := len(startTime) - 1; i >= 0; i-- {
		dp[i] = job[i][2]
		index := sort.SearchInts(startTime, job[i][1])
		if index < len(startTime) {
			dp[i] += dp[index]
		}
		if dp[i] < dp[i+1] {
			dp[i] = dp[i+1]
		}
	}
	return dp[0]
}
