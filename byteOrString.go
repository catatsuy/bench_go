package main

import "os"

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

	byteWrite(f, str)
	// stringWrite(f, str)
}

func byteWrite(f *os.File, str []byte) {
	defer f.Close()
	for i := 0; i < WRITE_COUNT; i++ {
		if _, err := f.Write(append(str, '\n')); err != nil {
			panic(err)
		}
	}
}

func stringWrite(f *os.File, str []byte) {
	defer f.Close()
	for i := 0; i < WRITE_COUNT; i++ {
		if _, err := f.WriteString(string(str) + "\n"); err != nil {
			panic(err)
		}
	}
}
