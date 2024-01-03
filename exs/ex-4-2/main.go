// Copyright © 2024 Victor N. Skurikhin.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
// Напишите программу, которая по умолчанию выводит дайджест
// SHA256 для входных данных, но при использовании соответствующих флагов
// командной строки выводит SHA384 или SHA512.

package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"os"
)

var stride = 1024

func ReadBytes(file *os.File) []byte {

	var result []byte
	buffer := make([]byte, 0, stride)
	reader := bufio.NewReader(file)
	for {
		n, err := io.ReadFull(reader, buffer[:cap(buffer)])
		buffer = buffer[:n]
		if err != nil {
			if err == io.EOF {
				break
			}
			if err != io.ErrUnexpectedEOF {
				//goland:noinspection GoUnhandledErrorResult
				fmt.Fprintln(os.Stderr, err)
				break
			}
		}
		result = append(result, buffer...)
	}
	return result
}

func readStdin() []byte {
	return ReadBytes(os.Stdin)
}

func main() {
	argsWithoutProg := os.Args[1:]
	b := readStdin()
	if len(argsWithoutProg) < 1 {
		fmt.Printf("%x\n", sha256.Sum256(b))
	} else {
		for _, s := range argsWithoutProg {
			switch s {
			case "-sha256":
				fmt.Printf("%x\n", sha256.Sum256(b))
				break
			case "-sha384":
				fmt.Printf("%x\n", sha512.Sum384(b))
				break
			case "-sha512":
				fmt.Printf("%x\n", sha512.Sum512(b))
				break
			}
		}
	}
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
