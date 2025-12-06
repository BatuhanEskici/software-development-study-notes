// Pattern: “Reverse Half of Number” (Half-Reversal Pattern)

// Time complexity = O(n)

// Space complexity = O(1)

// Solve version: 2 X
package main

func isPalindrome(x int) bool {
	if x < 0 || x > 0 && x%10 == 0 {
		return false
	}

	reversed := 0

	for x > reversed {
		lastDigit := x % 10

		reversed = reversed*10 + lastDigit

		x /= 10
	}

	return x == reversed || x == reversed/10
}
