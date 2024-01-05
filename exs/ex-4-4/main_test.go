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

func TestRotateLeft0(t *testing.T) {
	rotateLeftZeroSlice()
}

func TestRotateRight0(t *testing.T) {
	rotateRightZeroSlice()
}

func TestRotateLeft1(t *testing.T) {
	a := []int{0}
	RotateLeft(a, 1)
	gotool.AssertEqual(t, a[:], []int{0})
}

func TestRotateRight1(t *testing.T) {
	a := []int{0}
	RotateRight(a, 1)
	gotool.AssertEqual(t, a[:], []int{0})
}

func TestRotateLeft10(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateLeft(a, 0)
	gotool.AssertEqual(t, a[:], []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	RotateLeft(a, 1)
	gotool.AssertEqual(t, a[:], []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateLeft(a, 2)
	gotool.AssertEqual(t, a[:], []int{2, 3, 4, 5, 6, 7, 8, 9, 0, 1})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateLeft(a, 3)
	gotool.AssertEqual(t, a[:], []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateLeft(a, 4)
	gotool.AssertEqual(t, a[:], []int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateLeft(a, 5)
	gotool.AssertEqual(t, a[:], []int{5, 6, 7, 8, 9, 0, 1, 2, 3, 4})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateLeft(a, 6)
	gotool.AssertEqual(t, a[:], []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateLeft(a, 7)
	gotool.AssertEqual(t, a[:], []int{7, 8, 9, 0, 1, 2, 3, 4, 5, 6})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateLeft(a, 8)
	gotool.AssertEqual(t, a[:], []int{8, 9, 0, 1, 2, 3, 4, 5, 6, 7})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateLeft(a, 9)
	gotool.AssertEqual(t, a[:], []int{9, 0, 1, 2, 3, 4, 5, 6, 7, 8})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateLeft(a, 10)
	gotool.AssertEqual(t, a[:], []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateLeft(a, 11)
	gotool.AssertEqual(t, a[:], []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
}

func TestRotateRight10(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateRight(a, 0)
	gotool.AssertEqual(t, a[:], []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	RotateRight(a, 1)
	gotool.AssertEqual(t, a[:], []int{9, 0, 1, 2, 3, 4, 5, 6, 7, 8})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateRight(a, 2)
	gotool.AssertEqual(t, a[:], []int{8, 9, 0, 1, 2, 3, 4, 5, 6, 7})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateRight(a, 3)
	gotool.AssertEqual(t, a[:], []int{7, 8, 9, 0, 1, 2, 3, 4, 5, 6})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateRight(a, 4)
	gotool.AssertEqual(t, a[:], []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateRight(a, 5)
	gotool.AssertEqual(t, a[:], []int{5, 6, 7, 8, 9, 0, 1, 2, 3, 4})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateRight(a, 6)
	gotool.AssertEqual(t, a[:], []int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateRight(a, 7)
	gotool.AssertEqual(t, a[:], []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateRight(a, 8)
	gotool.AssertEqual(t, a[:], []int{2, 3, 4, 5, 6, 7, 8, 9, 0, 1})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateRight(a, 9)
	gotool.AssertEqual(t, a[:], []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateRight(a, 10)
	gotool.AssertEqual(t, a[:], []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	RotateRight(a, 11)
	gotool.AssertEqual(t, a[:], []int{9, 0, 1, 2, 3, 4, 5, 6, 7, 8})
}

func TestRotateLeftMaxIntAndAbove(t *testing.T) {
	rotateLeftMaxIntAndAbove([]int{0})
}
