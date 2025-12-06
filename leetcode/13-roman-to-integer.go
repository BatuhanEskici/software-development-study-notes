// Pattern: “Lookahead Pattern” or “Next-Value Comparison Pattern”

// Time complexity: O(n)

/* Space complexity: O(n)
Normally `runes := []rune(s)` uses O(n) time and O(n) extra memory.

But if using extra memory is due to the mandatory mechanisms
of the programming language and not to the problem-solving
logic, it is generally not included in space complexity.
*/

// Solve version: 2 ok
package main

var valuesOfSymbols = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var valueOfRoman int
	runes := []rune(s)

	for idx, runeValue := range runes {
		if idx+1 == len(runes) {
			valueOfRoman += valuesOfSymbols[runeValue]
			continue
		}

		if valuesOfSymbols[runeValue] < valuesOfSymbols[runes[idx+1]] {
			valueOfRoman -= valuesOfSymbols[runeValue]
		} else {
			valueOfRoman += valuesOfSymbols[runeValue]
		}
	}

	return valueOfRoman
}
