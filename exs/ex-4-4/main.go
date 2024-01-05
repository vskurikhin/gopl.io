// Copyright © 2024 Victor N. Skurikhin.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
// Напишите версию функции ro tate, которая работает в один
// проход.

package main

import (
	"fmt"
	"log"
	"math"
)

// Reverse reverses a slice of ints in place.
func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotateLeftNaive(s []int, r uint) {
	m := int(r) % len(s)
	if m != 0 {
		left := make([]int, m)
		for i := 0; i < m; i++ {
			left[i] = s[i]
		}
		for i := 0; (i + m) < len(s); i++ {
			s[i] = s[i+m]
		}
		for i, j := 0, len(s)-m; i < len(left) && j < len(s); i, j = i+1, j+1 {
			s[j] = left[i]
		}
	}
}

func RotateLeft(s []int, r uint) {
	m := int(r) % len(s)
	if m != 0 {
		left := make([]int, m)
		copy(left, s[:m])
		copy(s, s[m:])
		copy(s[len(s)-m:], left)
	}
}

func rotateRightNaive(s []int, r uint) {
	m := int(r) % len(s)
	if m != 0 {
		right := make([]int, m)
		for i, j := 0, len(s)-m; j < len(s); i, j = i+1, j+1 {
			right[i] = s[j]
		}
		for i := len(s) - 1; i >= m; i-- {
			s[i] = s[i-m]
		}
		for i, j := 0, 0; i < len(right); i, j = i+1, j+1 {
			s[i] = right[j]
		}
	}
}

func RotateRight(s []int, r uint) {
	m := int(r) % len(s)
	if m != 0 {
		right := make([]int, m)
		copy(right, s[len(s)-m:])
		for i := len(s) - 1; i >= m; i-- {
			s[i] = s[i-m]
		}
		copy(s, right)
	}
}

func rotateLeft2(s []int, r uint) {
	m := int(r) % len(s)
	if m != 0 {
		Reverse(s[:m])
		Reverse(s[m:])
		Reverse(s)
	}
}

func rotateRight2(s []int, r uint) {
	m := int(r) % len(s)
	if m != 0 {
		Reverse(s[m:])
		Reverse(s[:m])
		Reverse(s)
	}
}

func rotateLeftZeroSlice() {
	a := []int{}
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	RotateLeft(a, 0)
}

func rotateRightZeroSlice() {
	a := []int{}
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	RotateRight(a, 0)
}

func rotateLeftMaxIntAndAbove(s []int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	RotateLeft(s, uint(math.MaxInt)+2)
}

func main() {

	s := []int{0, 1, 2, 3}

	fmt.Printf("s = %v\n", s)
	RotateLeft(s, 0)
	fmt.Printf("RotateLeft(s, 0) s = %v\n", s)
	RotateLeft(s, 1)
	fmt.Printf("RotateLeft(s, 1) s = %v\n", s)
	RotateLeft(s, 2)
	fmt.Printf("RotateLeft(s, 2) s = %v\n", s)
	RotateLeft(s, 3)
	fmt.Printf("RotateLeft(s, 3) s = %v\n", s)
	RotateLeft(s, 4)
	fmt.Printf("RotateLeft(s, 4) s = %v\n", s)
	RotateLeft(s, 2)
	fmt.Printf("RotateLeft(s, 2) s = %v\n", s)
	RotateLeft(s, 2)
	fmt.Printf("RotateLeft(s, 2) s = %v\n", s)
	RotateLeft(s, 2)
	fmt.Printf("RotateLeft(s, 2) s = %v\n", s)
	RotateLeft(s, 4)
	fmt.Printf("RotateLeft(s, 4) s = %v\n", s)
	RotateLeft(s, 5)
	fmt.Printf("RotateLeft(s, 5) s = %v\n", s)
	RotateLeft(s, 6)
	fmt.Printf("RotateLeft(s, 6) s = %v\n", s)
	RotateLeft(s, 5)
	fmt.Printf("RotateLeft(s, 5) s = %v\n", s)

	rotateLeftMaxIntAndAbove(s)

	rotateRightNaive(s, 1)
	fmt.Printf("rotateRightNaive(s, 1) s = %v\n", s)
	rotateRightNaive(s, 2)
	fmt.Printf("rotateRightNaive(s, 2) s = %v\n", s)
	rotateRightNaive(s, 3)
	fmt.Printf("rotateRightNaive(s, 3) s = %v\n", s)
	rotateRightNaive(s, 4)
	fmt.Printf("rotateRightNaive(s, 4) s = %v\n", s)
	rotateRightNaive(s, 5)
	fmt.Printf("rotateRightNaive(s, 5) s = %v\n", s)
	rotateRightNaive(s, 1)
	fmt.Printf("rotateRightNaive(s, 1) s = %v\n", s)

	RotateRight(s, 1)
	fmt.Printf("RotateRight(s, 1) s = %v\n", s)
	RotateRight(s, 2)
	fmt.Printf("RotateRight(s, 2) s = %v\n", s)
	RotateRight(s, 3)
	fmt.Printf("RotateRight(s, 3) s = %v\n", s)
	RotateRight(s, 4)
	fmt.Printf("RotateRight(s, 4) s = %v\n", s)
	RotateRight(s, 5)
	fmt.Printf("RotateRight(s, 5) s = %v\n", s)
	RotateRight(s, 1)
	fmt.Printf("RotateRight(s, 1) s = %v\n", s)
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
