package main

import (
	"bufio"
	"os"
)

const (
	WRITE_COUNT = 10000000
)

func main() {
	str := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	fileName := "tmp.log"
	ferr := os.Remove(fileName)
	if ferr != nil {
		if !os.IsNotExist(ferr) {
			panic(ferr.Error())
		}
	}
	f, ferr := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if ferr != nil {
		panic(ferr.Error())
	}

	bufioWrite(f, str)
	// directWrite(f, str)
}

func bufioWrite(f *os.File, str []byte) {
	defer f.Close()
	w := bufio.NewWriterSize(f, 8192/2)
	defer w.Flush()
	for i := 0; i < WRITE_COUNT; i++ {
		if w.Available() <= len(str) {
			w.Flush()
		}
		if _, err := w.Write(append(str, '\n')); err != nil {
			panic(err)
		}
		w.Flush()
	}
}

func directWrite(f *os.File, str []byte) {
	defer f.Close()
	for i := 0; i < WRITE_COUNT; i++ {
		if _, err := f.Write(append(str, '\n')); err != nil {
			panic(err)
		}
	}
}
