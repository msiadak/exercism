// Package raindrops provides a function to convert integers to "raindrop speak"
package raindrops

import "strconv"

// Convert converts an integer to a string, the contents of which depend on the
// number's factors.
//
// - If the number has 3 as a factor, we add "Pling" to the output string.
// - If the number has 5 as a factor, we add "Plang".
// - If the number has 7 as a factor, we add "Plong".
// - If the number does not have 3, 5, or 7 as a factor, we just return the
// number converted to a string.
func Convert(n int) string {
	var output string
	if n%3 == 0 {
		output += "Pling"
	}
	if n%5 == 0 {
		output += "Plang"
	}
	if n%7 == 0 {
		output += "Plong"
	}
	if output == "" {
		output = strconv.Itoa(n)
	}
	return output
}
