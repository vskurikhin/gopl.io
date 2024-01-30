// Copyright © 2024 Victor N. Skurikhin.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 129.
// Напишите программу wordfreq для подсчета частоты каждого
// слова во входном текстовом файле. Вызовите input.Split(bufio.ScanWords) до
// первого вызова Scan для разбивки текста на слова, а не на строки.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WordFreq(reader io.Reader) map[string]int {

	freq := make(map[string]int)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		buffer := scanner.Bytes()
		freq[string(buffer)]++
	}
	return freq
}

func main() {

	fmt.Printf("word\tfreq\n")
	freq := WordFreq(os.Stdin)
	for c, n := range freq {
		fmt.Printf("%s\t\t%d\n", c, n)
	}
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
