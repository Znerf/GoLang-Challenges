
package main

import (
	"fmt"
	"sort"
	"strings"
)

func mix(s1, s2 string) string {
	s1Counts := make(map[rune]int)
	s2Counts := make(map[rune]int)

	// Count lowercase letter frequencies in s1
	for _, char := range s1 {
		if char >= 'a' && char <= 'z' {
			s1Counts[char]++
		}
	}

	// Count lowercase letter frequencies in s2
	for _, char := range s2 {
		if char >= 'a' && char <= 'z' {
			s2Counts[char]++
		}
	}

	var result strings.Builder

	// Iterate through lowercase letters in s1 and s2
	for char := 'a'; char <= 'z'; char++ {
		maxCount := max(s1Counts[char], s2Counts[char])

		// Only consider letters with maximum count > 1
		if maxCount > 1 {
			var prefix string
			if s1Counts[char] > 0 && s2Counts[char] > 0 {
				// If max count in both s1 and s2, use =: prefix
				prefix = "=:"
			} else if s1Counts[char] > 0 {
				// If max count only in s1, use 1: prefix
				prefix = "1:"
			} else {
				// If max count only in s2, use 2: prefix
				prefix = "2:"
			}

			// Append to result string
			result.WriteString(fmt.Sprintf("%s%s/", prefix, strings.Repeat(string(char), maxCount)))
		}
	}

	// Remove trailing '/' from result
	resultString := result.String()
	resultString = strings.TrimSuffix(resultString, "/")

	// Sort result by length in decreasing order, and by lexicographic order within the same length
	resultSlice := strings.Split(resultString, "/")
	sort.Slice(resultSlice, func(i, j int) bool {
		if len(resultSlice[i]) != len(resultSlice[j]) {
			return len(resultSlice[i]) > len(resultSlice[j])
		}
		return resultSlice[i] < resultSlice[j]
	})

	return strings.Join(resultSlice, "/")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test with examples
	s1 := "my&friend&Paul has heavy hats! &"
	s2 := "my friend John has many many friends &"
	fmt.Println(mix(s1, s2)) // Expected output: "2:nnnnn/1:aaaa/1:hhh/2:mmm/2:yyy/2:dd/2:ff/2:ii/2:rr/=:ee/=:ss"

	s1 = "mmmmm m nnnnn y&friend&Paul has heavy hats! &"
	s2 = "my frie n d Joh n has ma n y ma n y frie n ds n&"
	fmt.Println(mix(s1, s2)) // Expected output: "1:mmmmmm/=:nnnnnn/1:aaaa/1:hhh/2:yyy/2:dd/2:ff/2:ii/2:rr/=:ee/=:ss"

	s1 = "Are the kids at home? aaaaa fffff"
	s2 = "Yes they are here! aaaaa fffff"
	fmt.Println(mix(s1, s2)) // Expected output: "=:aaaaaa/
