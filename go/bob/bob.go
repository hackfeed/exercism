// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	var answer string
	isUpperCase := true
	isNonLetters := false
	isSilence
	for _, r := range remark {
		if unicode.IsUpper(r) {
			isUpperCase = false
		}
	}

	if remark == "" {
		answer = "Fine. Be that way!"
	} else {
		if isUpperCase == true {
			if remark[len(remark)-1:] == "?" {
				answer = "Calm down, I know what I'm doing!"
			} else {
				answer = "Whoa, chill out!"
			}
		} else {
			if remark[len(remark)-1:] == "?" {
				answer = "Sure."
			} else {
				answer = "Whatever."
			}
		}
	}

	return answer
}
