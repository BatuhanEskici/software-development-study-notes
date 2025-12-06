// Pattern: Vertical Scanning

// Time complexity: O(n*m)

// Space complexity: O(1)

// 2 ok

package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]

		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) || c != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
