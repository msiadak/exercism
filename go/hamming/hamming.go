// Package hamming provides a function for comparing strings representing DNA
// strands.
package hamming

import "errors"

// Distance returns the Hamming distance between two strands of DNA represented
// as strings.
//
// Doesn't check that the strings contain DNA nucleic acids, just compares them
// to see if the runes are the same.
//
// Returns -1 and an error if the input strings are not the same length.
func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		return -1, errors.New("input strings must be the same length")
	}
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, err
}
