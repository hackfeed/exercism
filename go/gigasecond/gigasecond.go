// Package gigasecond have functions to work with gigaseconds.
package gigasecond

import "time"

// AddGigasecond adds one gigasecond to passed date.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
