// Copyright © 2024 Victor N. Skurikhin.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
// Перепишите функцию reverse так, чтобы вместо среза она использовала
// указатель на массив.

package main

import "fmt"

func reverseArrayPtr1(s *[1]int) {
}

func reverseArrayPtr2(s *[2]int) {
	s[0], s[1] = s[1], s[0]
}

func reverseArrayPtr3(s *[3]int) {
	s[0], s[2] = s[2], s[0]
}

func reverseArrayPtr4(s *[4]int) {
	s[0], s[1], s[2], s[3] = s[3], s[2], s[1], s[0]
}

func reverseArrayPtr5(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseArrayPtr6(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseArrayPtr7(s *[7]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseArrayPtr8(s *[8]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseArrayPtr9(s *[9]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseSicePtr(p *[]int) {
	s := *p
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0}
	fmt.Printf("a = %v\n", a)
	reverseArrayPtr1(&a)
	fmt.Printf("a = %v\n", a)
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
