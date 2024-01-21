// Copyright © 2024 Victor N. Skurikhin.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
// Перепишите функцию reverse так, чтобы она без выделения
// дополнительной памяти обращала последовательность символов среза []byte, который
// представляет строку в кодировке UTF-8. Сможете ли вы обойтись без выделения
// новой памяти?

package main

import (
	"fmt"
	"unicode/utf8"
)

func Print(buffer []byte) {
	fmt.Printf(
		"len(%s) = %d, utf8.RuneCount(%s) = %d\n",
		string(buffer), len(buffer), string(buffer), utf8.RuneCount(buffer),
	)

	for _, r := range string(buffer) {
		fmt.Printf("utf8.RuneLen(%c) = %d\n", r, utf8.RuneLen(r))
	}
}

func writeRuneAt(buffer *[]byte, current rune, i int) {
	pb := (*buffer)[i:i]
	utf8.AppendRune(pb, current)
}

// Reverse reverses a slice of ints in place.
func Reverse(buffer *[]byte) {

	if utf8.RuneCount(*buffer) < 2 {
		return
	}
	left, widthLeft := utf8.DecodeRune(*buffer)
	right, widthRight := utf8.DecodeLastRune(*buffer)

	switch utf8.RuneCount(*buffer) {
	case 2:
		writeRuneAt(buffer, right, 0)
		writeRuneAt(buffer, left, widthRight)
	case 3:
		center, widthCenter := utf8.DecodeRune((*buffer)[widthLeft:])
		writeRuneAt(buffer, right, 0)
		writeRuneAt(buffer, center, widthRight)
		writeRuneAt(buffer, left, widthRight+widthCenter)
	default:
		pb := (*buffer)[widthLeft : len(*buffer)-widthRight]
		Reverse(&pb)
		copy((*buffer)[widthRight:], pb)
		writeRuneAt(buffer, right, 0)
		writeRuneAt(buffer, left, widthRight+len(pb))
	}
}

//goland:noinspection GoSnakeCaseUsage
func main() {
	// utf8.RuneCount == 4
	var t1 = []byte("1234")
	Reverse(&t1)
	fmt.Printf("Reverse(t1) = %s\n\n", string(t1))

	var t2 = []byte("1²³⁴")
	Print(t2)
	Reverse(&t2)
	fmt.Printf("Reverse(t2) = %s\n\n\n", string(t2))
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
