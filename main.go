// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "fmt"

type Pair struct {
	a, b int
}

func main() {
	nums := []int{2, 4, 3, 5, 7, 8, 9, 2, 4, 3}
	target := 10
	pairs := make(map[int]Pair)

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			a := nums[i]
			b := nums[j]
			if a > b {
				temp := a
				a = b
				b = temp
			}
			if a+b == target {
				pairs[a] = Pair{a, b}
			}
		}
	}

	for _, pair := range pairs {
		fmt.Println(pair.a, pair.b)
	}

}

//Algorithm Challenge (Sorting and Logic): Write a Python function to find all
// pairs in a list that sum up to a given target.
//Example:
//Input: nums = [2, 4, 3, 5, 7, 8, 9], target = 10
//Output: [(3, 7), (2, 8)]
