// Package acronym flex by it's own.
package acronym

import "strings"

// Abbreviate make an abbreviation of given string.
func Abbreviate(s string) string {
	abb := ""
	r := strings.NewReplacer("'", "", "_", "", "-", " ")
	s = r.Replace(s)
	words := strings.Split(s, " ")
	for _, word := range words {
		if len(word) != 0 {
			abb += string(strings.Title(word)[0])
		}
	}

	return abb
}
