package main

import (
	"bufio"
	"os"
)

func bufioWrite(w *bufio.Writer, f *os.File, str []byte) {
	defer w.Flush()
	if w.Available() <= len(str) {
		w.Flush()
	}
	if _, err := w.Write(append(str, '\n')); err != nil {
		panic(err)
	}
	w.Flush()
}

func directWrite(f *os.File, str []byte) {
	if _, err := f.Write(append(str, '\n')); err != nil {
		panic(err)
	}
}
