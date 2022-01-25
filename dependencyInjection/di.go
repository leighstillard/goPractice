package main

import (
	"io"
	"os"
)

// It returns the number of bytes writte and any write error encountered.
func Greet(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error} (
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
))

type Writer interface {
	Write(p []byte) (n int, err error)
}