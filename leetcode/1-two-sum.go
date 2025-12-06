// Pattern: Hash Map (Dictionary) Complement Lookup Pattern

// Time complexity: O(n)

// Space complexity: O(n)

// Solve version: 2 ok
package main

func twoSum(nums []int, target int) []int {
	numberMap := map[int]int{}

	for i, val := range nums {
		valueNeeded := target - val

		if _, ok := numberMap[valueNeeded]; ok {
			return []int{numberMap[valueNeeded], i}
		}

		numberMap[val] = i
	}

	return []int{}
}
