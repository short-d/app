package mdio

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"

	"github.com/short-d/app/fw"
)

var _ fw.StdIn = (*StdIn)(nil)

type StdIn struct {
}

func (s StdIn) Read(p []byte) (n int, err error) {
	return os.Stdin.Read(p)
}

func NewBuildInStdIn() StdIn {
	return StdIn{}
}

var _ fw.StdOut = (*StdOut)(nil)

type StdOut struct {
}

func (s StdOut) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

func NewBuildInStdOut() StdOut {
	return StdOut{}
}

func Tap(r io.ReadCloser, fn func(text string)) io.ReadCloser {
	buf, _ := ioutil.ReadAll(r)
	text := string(buf)

	fn(text)

	return ioutil.NopCloser(bytes.NewBuffer(buf))
}
