// Package twofer provides a function to assist children with sharing.
package twofer

import "fmt"

// ShareWith returns the string "One for {name}, one for me."
// When name is not given, assume "you" is the name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
