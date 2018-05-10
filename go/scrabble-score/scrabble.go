// Package scrabble provides a function for calculating scrabble scores.
package scrabble

import (
	"regexp"
	"strings"
)

var patternScores = map[string]int{
	"[aeioulnrst]": 1,
	"[dg]":         2,
	"[bcmp]":       3,
	"[fhvwy]":      4,
	"[k]":          5,
	"[jx]":         8,
	"[qz]":         10,
}

// Score accepts a string (a scrabble word) and returns the Scrabble score
// for that word.
//
// Score key:
// A, E, I, O, U, L, N, R, S, T:  1
// D, G:									        2
// B, C, M, P:                    3
// F, H, V, W, Y:                 4
// K:                             5
// J, X:                          8
// Q, Z:                         10
func Score(word string) int {
	var score int
	for _, r := range strings.ToLower(word) {
		for pattern, letterScore := range patternScores {
			re := regexp.MustCompile(pattern)
			if re.MatchString(string(r)) {
				score += letterScore
			}
		}
	}
	return score
}
