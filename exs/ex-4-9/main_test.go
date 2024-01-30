package main

import (
	"flag"
	"github.com/vskurikhin/gotool"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestWithoutArgs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"ex-4-8"}
	main()
}

// !+test
func TestWordFreq(t *testing.T) {
	var emptyMap = make(map[string]int)
	var tests = []struct {
		input string
		want  map[string]int
	}{
		{"", emptyMap},
		{"a", map[string]int{"a": 1}},
		{"a a", map[string]int{"a": 2}},
		{"a a b", map[string]int{"a": 2, "b": 1}},
		{
			`// Copyright © 2024 Victor N. Skurikhin. // License: https://creativecommons.org/licenses/by-nc-sa/4.0/ // See page 129. // Напишите программу wordfreq для подсчета частоты каждого // слова во входном текстовом файле. Вызовите input.Split(bufio.ScanWords) до // первого вызова Scan для разбивки текста на слова, а не на строки. package main import ( "bufio" "fmt" "io" "os" ) func WordFreq(reader io.Reader) map[string]int { freq := make(map[string]int) scanner := bufio.NewScanner(reader) scanner.Split(bufio.ScanWords) for scanner.Scan() { buffer := scanner.Bytes() freq[string(buffer)]++ } return freq } func main() { fmt.Printf("word\tfreq\n") freq := WordFreq(os.Stdin) for c, n := range freq { fmt.Printf("%s\t\t%d\n", c, n) } } //!- /* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */`,
			map[string]int{
				"https://creativecommons.org/licenses/by-nc-sa/4.0/": 1,
				"во":                             1,
				")":                              1,
				"Skurikhin.":                     1,
				"License:":                       1,
				"разбивки":                       1,
				"package":                        1,
				"func":                           2,
				"bufio.NewScanner(reader)":       1,
				"main()":                         1,
				"n)":                             1,
				"See":                            1,
				"Вызовите":                       1,
				"wordfreq":                       1,
				"входном":                        1,
				"слова,":                         1,
				"import":                         1,
				"freq":                           4,
				"range":                          1,
				"N.":                             1,
				"page":                           1,
				"io.Reader)":                     1,
				"/*":                             1,
				"частоты":                        1,
				"(":                              1,
				"не":                             1,
				"строки.":                        1,
				`"fmt"`:                          1,
				"c,":                             2,
				"vim:":                           1,
				"tabstop=4":                      1,
				"Copyright":                      1,
				"текстовом":                      1,
				`"os"`:                           1,
				"softtabstop=4":                  1,
				`"io"`:                           1,
				"scanner.Split(bufio.ScanWords)": 1,
				"scanner.Scan()":                 1,
				"n":                              1,
				"Victor":                         1,
				"каждого":                        1,
				"WordFreq(os.Stdin)":             1,
				"первого":                        1,
				"а":                              1,
				"map[string]int":                 1,
				":=":                             5,
				`fmt.Printf("%s\t\t%d\n",`:       1,
				"*/":                             1,
				"129.":                           1,
				"input.Split(bufio.ScanWords)":   1,
				"scanner":                        1,
				"}":                              4,
				"shiftwidth=4":                   1,
				"для":                            2,
				"на":                             2,
				"scanner.Bytes()":                1,
				"return":                         1,
				"//":                             6,
				"buffer":                         1,
				`"bufio"`:                        1,
				"for":                            2,
				"Напишите":                       1,
				"до":                             1,
				"программу":                      1,
				"файле.":                         1,
				"текста":                         1,
				"WordFreq(reader":                1,
				"//!-":                           1,
				"слова":                          1,
				"Scan":                           1,
				"main":                           1,
				"{":                              4,
				"freq[string(buffer)]++":         1,
				`fmt.Printf("word\tfreq\n")`:     1,
				"noexpandtab:":                   1,
				"©":                              1,
				"подсчета":                       1,
				"make(map[string]int)":           1,
				"set":                            1,
				"2024":                           1,
				"вызова":                         1,
			},
		},
	}
	for _, test := range tests {
		reader := strings.NewReader(test.input)
		m := WordFreq(reader)
		gotool.AssertEqual(t, m, test.want)
	}
}

//!-bench
