// Package gigasecond exports a single function (AddGigasecond) that performs
// a trivial time calculation.
package gigasecond

import "time"

// AddGigasecond adds 10e9 seconds to the input time and returns the result.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000 * time.Second)
}
