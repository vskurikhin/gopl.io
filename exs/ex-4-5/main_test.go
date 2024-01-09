package main

import (
	"flag"
	"github.com/vskurikhin/gotool"
	"os"
	"testing"
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
	os.Args = []string{"ex-4-4"}
	main()
}

// !+test
func TestDeduplicate(t *testing.T) {
	var tests = []struct {
		input []string
		want  []string
	}{
		{nil, nil},
		{[]string{}, []string{}},
		{[]string{"a"}, []string{"a"}},
		{[]string{"a", "a"}, []string{"a"}},
		{[]string{"a", "b"}, []string{"a", "b"}},
		{[]string{"a", "a", "b"}, []string{"a", "b"}},
		{[]string{"a", "b", "b"}, []string{"a", "b"}},
		{[]string{"a", "b", "b", "c"}, []string{"a", "b", "c"}},
		{[]string{"a", "a", "b", "b", "c", "c"}, []string{"a", "b", "c"}},
	}
	for _, test := range tests {
		got := Deduplicate(test.input)
		gotool.AssertEqual(t, got, test.want)
	}
}

//!-test

// !+bench
func BenchmarkDeduplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Deduplicate([]string{"a", "a", "b", "b", "c", "c"})
	}
}

//!-bench
