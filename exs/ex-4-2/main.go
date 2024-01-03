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

var stride = 4096

//goland:noinspection GoUnhandledErrorResult
var PrintlnStderr = func(err error) { fmt.Fprintln(os.Stderr, err) }

func IsErrEOF(err error, unknownErrHook func()) bool {

	switch {
	case err == nil:
		return false
	case err == io.EOF:
		return true
	case err != io.ErrUnexpectedEOF:
		unknownErrHook()
		return true
	}
	return false
}

func ReadBytes(file *os.File) []byte {

	buffer := make([]byte, 0, stride)
	reader := bufio.NewReader(file)
	result := make([]byte, 0, stride)
	for {
		n, err := io.ReadFull(reader, buffer[:cap(buffer)])
		buffer = buffer[:n]
		if IsErrEOF(err, func() { PrintlnStderr(err) }) {
			break
		}
		result = append(result, buffer...)
	}
	return result
}

func readStdin() []byte {
	return ReadBytes(os.Stdin)
}

func caseByArgs(argsWithoutProg []string, buffer []byte) {

	for _, s := range argsWithoutProg {
		switch s {
		default:
			fmt.Printf("%x\n", sha256.Sum256(buffer))
		case "-sha384":
			fmt.Printf("%x\n", sha512.Sum384(buffer))
		case "-sha512":
			fmt.Printf("%x\n", sha512.Sum512(buffer))
		}
	}
}

func main() {

	argsWithoutProg := os.Args[1:]
	buffer := readStdin()
	if len(argsWithoutProg) < 1 {
		fmt.Printf("%x\n", sha256.Sum256(buffer))
	} else {
		caseByArgs(argsWithoutProg, buffer)
	}
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
