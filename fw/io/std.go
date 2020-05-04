package io

import (
	"os"
)

var _ Input = (*StdIn)(nil)

type StdIn struct {
}

func (s StdIn) Read(p []byte) (n int, err error) {
	return os.Stdin.Read(p)
}

func NewStdIn() StdIn {
	return StdIn{}
}

var _ Output = (*StdOut)(nil)

type StdOut struct {
}

func (s StdOut) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

func NewStdOut() StdOut {
	return StdOut{}
}
