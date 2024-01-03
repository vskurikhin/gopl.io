package main

import (
	"flag"
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
	os.Args = []string{"ex-4-2"}
	main()
}

func TestArgsSHA256(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"ex-4-2", "-sha256"}
	main()
}

func TestArgsSHA384(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"ex-4-2", "-sha384"}
	main()
}

func TestArgsSHA512(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"ex-4-2", "-sha512"}
	main()
}
