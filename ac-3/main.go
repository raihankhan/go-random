package main

import (
	"bufio"
	"fmt"
	"os"
)

func maxSubarrayLength(arr []int, n, S int) int {
	left, sum, maxLength := 0, 0, 0

	for right := 0; right < n; right++ {
		// Add the current element to the sum
		sum += arr[right]

		// Shrink the window from the left if the sum exceeds S
		for sum > S && left <= right {
			sum -= arr[left]
			left++
		}

		// Update the maximum length of the valid subarray
		if sum <= S {
			maxLength = max(maxLength, right-left+1)
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, S int
	fmt.Fscan(reader, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	fmt.Fscan(reader, &S)

	result := maxSubarrayLength(arr, n, S)
	fmt.Println(result)
}
