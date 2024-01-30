package main

import (
	"flag"
	"github.com/vskurikhin/gotool"
	"os"
	"os/exec"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestMain(m *testing.M) {
	flag.Parse()
	exitCode := m.Run()
	// Exit
	os.Exit(exitCode)
}

func TestWithoutArgs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"ex-4-8"}
	main()
}

// !+test
func TestReverse(t *testing.T) {
	var emptyMap = make(map[string]int)
	var zeroSlice = make([]int, utf8.UTFMax+1)
	var tests = []struct {
		input string
		want1 map[string]int
		want2 []int
	}{
		{"", emptyMap, zeroSlice},
		{"a", map[string]int{Letter.String(): 1}, []int{0, 1, 0, 0, 0}},
		{
			`poekgpjem'vpjm905g ijg 4805u 04tj-20i349058839^$&*&*(@@)*#_)(*(*&(*^
зуок0пего4хщшопжмь лцд.о45они  ькдеь0495686988085?%((*"_*"_;"+;)%(*:
`,
			map[string]int{
				Control.String(): 2,
				Letter.String():  49,
				Number.String():  37,
				Space.String():   6,
				Symbol.String():  4,
			},
			[]int{0, 110, 28, 0, 0},
		},
	}
	for _, test := range tests {
		reader := strings.NewReader(test.input)
		m, s := CharCount(reader)
		gotool.AssertEqual(t, m, test.want1)
		gotool.AssertEqual(t, s, test.want2)
	}
}

func TestInput(t *testing.T) {
	subProc := exec.Command("cat")
	input := "abc\n"
	subProc.Stdin = strings.NewReader(input)
	output, _ := subProc.Output()

	if input != string(output) {
		t.Errorf("Wanted: %v, Got: %v", input, string(output))
	}
	subProc.Wait()
}

//!-bench
