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
	os.Args = []string{"ex-4-3"}
	main()
}

func TestReverse1(t *testing.T) {
	a := [...]int{0}
	reverseArrayPtr1(&a)
	gotool.AssertEqual(t, a[:], []int{0})
}

func TestReverse2(t *testing.T) {
	a := [...]int{0, 1}
	reverseArrayPtr2(&a)
	gotool.AssertEqual(t, a[:], []int{1, 0})
}

func TestReverse3(t *testing.T) {
	a := [...]int{0, 1, 3}
	reverseArrayPtr3(&a)
	gotool.AssertEqual(t, a[:], []int{3, 1, 0})
}

func TestReverse4(t *testing.T) {
	a := [...]int{0, 1, 3, 4}
	reverseArrayPtr4(&a)
	gotool.AssertEqual(t, a[:], []int{4, 3, 1, 0})
}

func TestReverse5(t *testing.T) {
	a := [...]int{0, 1, 3, 4, 5}
	reverseArrayPtr5(&a)
	gotool.AssertEqual(t, a[:], []int{5, 4, 3, 1, 0})
}

func TestReverse6(t *testing.T) {
	a := [...]int{0, 1, 3, 4, 5, 6}
	reverseArrayPtr6(&a)
	gotool.AssertEqual(t, a[:], []int{6, 5, 4, 3, 1, 0})
}

func TestReverse7(t *testing.T) {
	a := [...]int{0, 1, 3, 4, 5, 6, 7}
	reverseArrayPtr7(&a)
	gotool.AssertEqual(t, a[:], []int{7, 6, 5, 4, 3, 1, 0})
}

func TestReverse8(t *testing.T) {
	a := [...]int{0, 1, 3, 4, 5, 6, 7, 8}
	reverseArrayPtr8(&a)
	gotool.AssertEqual(t, a[:], []int{8, 7, 6, 5, 4, 3, 1, 0})
}

func TestReverse9(t *testing.T) {
	a := [...]int{0, 1, 3, 4, 5, 6, 7, 8, 9}
	reverseArrayPtr9(&a)
	gotool.AssertEqual(t, a[:], []int{9, 8, 7, 6, 5, 4, 3, 1, 0})
}

func TestReverseSlice(t *testing.T) {
	a := []int{0}
	reverseSicePtr(&a)
	gotool.AssertEqual(t, a, []int{0})
	b := []int{0, 1}
	reverseSicePtr(&b)
	gotool.AssertEqual(t, b, []int{1, 0})
	c := []int{0, 1, 3}
	reverseSicePtr(&c)
	gotool.AssertEqual(t, c, []int{3, 1, 0})
}
