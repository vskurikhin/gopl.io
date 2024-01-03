// Copyright © 2024 Victor N. Skurikhin.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 303.
//!+

// Package word provides utilities for word games.
package word

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func letterStringsReBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if unicode.IsLetter(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func IsEqualRunes(x, y rune) bool {
	return unicode.ToLower(x) == unicode.ToLower(y)
}

func at(pstr *string, index int) rune {
	runes := []rune(*pstr)
	return runes[index]
}

// IsPalindrome reports whether s reads the same forward and backward.
// (Our first attempt.)
func IsPalindrome(str string) bool {
	s := letterStringsReBuilder(str)
	n := utf8.RuneCountInString(s)
	for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
		r1, r2 := at(&s, i), at(&s, j)
		if !IsEqualRunes(r1, r2) {
			return false
		}
	}
	return true
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
