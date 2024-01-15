// Copyright © 2024 Victor N. Skurikhin.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
// Напишите функцию, которая без выделения дополнительной памяти преобразует
// последовательности смежных пробельных символов Unicode
// (см. Unicode.IsSpace ) в срезе []byte в кодировке UTF-8 в один пробел ASCII.

package main

import (
	"fmt"
	"github.com/vskurikhin/gotool"
	"strings"
	"unicode"
	"unicode/utf8"
)

func writeRuneStringBuilder(sb *strings.Builder, r rune) {
	if unicode.IsSpace(r) {
		sb.WriteRune(' ')
	} else {
		sb.WriteRune(r)
	}
}

func TrimSpaceString(str string) string {

	var sb strings.Builder
	previous, width := utf8.DecodeRuneInString(str)
	writeRuneStringBuilder(&sb, previous)
	for i := width; i < len(str); i += width {
		current, w := utf8.DecodeRuneInString(str[i:])
		if !unicode.IsSpace(previous) || !unicode.IsSpace(current) {
			writeRuneStringBuilder(&sb, current)
		}
		previous = current
		width = w
	}
	return sb.String()
}

func writeRune(buffer *[]byte, current rune, c *int) {
	if unicode.IsSpace(current) {
		(*buffer)[*c] = ' '
		*c = *c + 1
	} else {
		pb := (*buffer)[*c:*c]
		*c = *c + utf8.RuneLen(current)
		utf8.AppendRune(pb, current)
	}
}

func TrimSpace(buffer *[]byte) *[]byte {

	if len(*buffer) < 1 {
		return buffer
	}
	var c = 0
	previous, width := utf8.DecodeRune(*buffer)
	writeRune(buffer, previous, &c)
	for i := width; i < len(*buffer) && c < len(*buffer); i += width {
		current, w := utf8.DecodeRune((*buffer)[i:])
		if !unicode.IsSpace(previous) || !unicode.IsSpace(current) {
			writeRune(buffer, current, &c)
		}
		previous = current
		width = w
	}
	result := (*buffer)[:c]
	return &result
}

func main() {
	var b = []byte("a  a   b b    c c")
	got1 := TrimSpace(&b)
	fmt.Printf("%s\n", string(*got1))
	gotool.AssertEqual(nil, []byte("a a b b c c"), *got1)
	got2 := TrimSpaceString("a  a   b b    c c")
	if got2 != "a a b b c c" {
		panic(`IsPalindrome("palindrome") = true`)
	}
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
