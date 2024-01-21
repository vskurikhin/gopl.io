package main

import (
	"flag"
	"fmt"
	"github.com/vskurikhin/gotool"
	"os"
	"testing"
	"unicode/utf8"
)

var buffer []byte
var original []rune
var str string

func init() {
	original = []rune{'1', '2', '3', '4', '¹', '²', '³', '⁴', '①', '②', '③', '④'}
	str = string(original)
	buffer = []byte(str)
}

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
func TestReverse(t *testing.T) {
	var tests = []struct {
		input []byte
		want  []byte
	}{
		{nil, nil},
		{[]byte{}, []byte{}},
		{[]byte(""), []byte("")},
		{[]byte("1"), []byte("1")},
		{[]byte("12"), []byte("21")},
		{[]byte("1²"), []byte("²1")},
		{[]byte("¹2"), []byte("2¹")},
		{[]byte("1②"), []byte("②1")},
		{[]byte("①2"), []byte("2①")},
		{[]byte("¹²"), []byte("²¹")},
		{[]byte("123"), []byte("321")},
		{[]byte("12³"), []byte("³21")},
		{[]byte("1²3"), []byte("3²1")},
		{[]byte("¹23"), []byte("32¹")},
		{[]byte("12③"), []byte("③21")},
		{[]byte("1②3"), []byte("3②1")},
		{[]byte("①23"), []byte("32①")},
		{[]byte("1²③"), []byte("③²1")},
		{[]byte("1②³"), []byte("³②1")},
		{[]byte("¹2③"), []byte("③2¹")},
		{[]byte("¹②3"), []byte("3②¹")},
		{[]byte("①2³"), []byte("³2①")},
		{[]byte("①²3"), []byte("3²①")},
	}
	for _, test := range tests {
		var i = test.input
		var w = test.want
		Reverse(&i)
		gotool.AssertEqual(t, i, w)
	}
}

func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func TestReversePermutations(t *testing.T) {
	permutations := Permutations(original, 4)
	toSlice, _ := permutations.Slice()
	fmt.Println("Total:", permutations.Count())
	for _, elem := range toSlice {
		str := string(elem)
		rev := reverse(str)
		buffer := []byte(string(elem))
		Reverse(&buffer)
		got := string(buffer)
		if got != rev {
			t.Fatalf("got = %s, rev = %s\n", got, rev)
		}
	}
}

//!-test

// !+bench
func BenchmarkMyReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse(&buffer)
	}
}

func BenchmarkCopyReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse(str)
	}
}

//!-bench
