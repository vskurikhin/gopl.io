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
func TestTrimSpace(t *testing.T) {
	var tests = []struct {
		input []byte
		want  []byte
	}{
		{nil, nil},
		{[]byte{}, []byte{}},
		{[]byte("a"), []byte("a")},
		{[]byte(" a"), []byte(" a")},
		{[]byte("  a"), []byte(" a")},
		{[]byte("   a"), []byte(" a")},
		{[]byte("a "), []byte("a ")},
		{[]byte("a  "), []byte("a ")},
		{[]byte("a   "), []byte("a ")},
		{[]byte("a  a   b b    c c"), []byte("a a b b c c")},
	}
	for _, test := range tests {
		var i = test.input
		var w = test.want
		got := TrimSpace(&i)
		gotool.AssertEqual(t, got, &w)
	}
}

//!-test

// !+bench
func BenchmarkTrimSpace(b *testing.B) {
	var buffer = []byte("a  a   b b    c c")
	for i := 0; i < b.N; i++ {
		TrimSpace(&buffer)
	}
}

func BenchmarkTrimSpaceString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrimSpaceString("a a b b c c")
	}
}

//!-bench
