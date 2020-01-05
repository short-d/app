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
	reader io.Reader
}

func (s StdIn) Read(p []byte) (n int, err error) {
	return s.reader.Read(p)
}

func NewBuildInStdIn() StdIn {
	return StdIn{os.Stdin}
}

var _ fw.StdOut = (*StdOut)(nil)

type StdOut struct {
	writer io.Writer
}

func (s StdOut) Write(p []byte) (n int, err error) {
	return s.writer.Write(p)
}

func NewBuildInStdOut() StdOut {
	return StdOut{writer: os.Stdout}
}

func Tap(r io.ReadCloser, fn func(text string)) io.ReadCloser {
	buf, _ := ioutil.ReadAll(r)
	text := string(buf)

	fn(text)

	return ioutil.NopCloser(bytes.NewBuffer(buf))
}
