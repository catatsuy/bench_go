package main

import "os"

func byteWrite(f *os.File, str []byte) {
	if _, err := f.Write(append(str, '\n')); err != nil {
		panic(err)
	}
}

func stringWrite(f *os.File, str []byte) {
	if _, err := f.WriteString(string(str) + "\n"); err != nil {
		panic(err)
	}
}
