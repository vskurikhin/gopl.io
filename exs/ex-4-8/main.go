// Copyright © 2024 Victor N. Skurikhin.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 129.
// Измените сhaгcount так, чтобы программа подсчитывала количество букв,
// цифр и прочих категорий Unicode с использованием функций наподобие
// Unicode.IsLetter.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type Сategory int

const (
	Control Сategory = iota
	Letter
	Mark
	Number
	Space
	Symbol
)

func (d Сategory) String() string {
	return [...]string{"Control", "Letter", "Mark", "Number", "Space", "Symbol"}[d]
}

func main() {
	counts := make(map[string]int)  // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		switch {
		case unicode.IsControl(r):
			counts[Control.String()]++
		case unicode.IsLetter(r):
			counts[Letter.String()]++
		case unicode.IsMark(r):
			counts[Mark.String()]++
		case unicode.IsNumber(r):
			counts[Number.String()]++
		case unicode.IsSpace(r):
			counts[Space.String()]++
		case unicode.IsSymbol(r):
			counts[Symbol.String()]++
		}
		utflen[n]++
	}
	fmt.Printf("Сategory\tcount\n")
	for c, n := range counts {
		fmt.Printf("%s\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
