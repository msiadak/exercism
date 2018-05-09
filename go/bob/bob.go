// Package bob provides a function that simulates speaking with a lackadasical
// teenager named Bob. Bob responds with one of five responses to any remark.
package bob

import "regexp"

// Hey accepts a remark string and returns Bob's creative response to that
// remark.
func Hey(remark string) string {
	switch {
	case mustMatchString(`^([^a-z]*[A-Z]+[^a-z]*)+\?\s*$`, remark):
		return "Calm down, I know what I'm doing!"
	case mustMatchString(`^([^a-z]*[A-Z]+[^a-z]*)+$`, remark):
		return "Whoa, chill out!"
	case mustMatchString(`^.*\?\s*$`, remark):
		return "Sure."
	case mustMatchString(`^\s*$`, remark):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}

// mustMatchString returns whether s matches pattern and panics if
// MatchString returns an error.
func mustMatchString(pattern string, s string) bool {
	match, err := regexp.MatchString(pattern, s)
	if err != nil {
		panic(err)
	}
	return match
}
