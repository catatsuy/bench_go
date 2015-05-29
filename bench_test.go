package main

import (
	"bufio"
	"os"
	"testing"
)

func BenchmarkCloseEvery(b *testing.B) {
	str := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	fileName := "tmp.log"
	ferr := os.Remove(fileName)
	if ferr != nil {
		if !os.IsNotExist(ferr) {
			panic(ferr.Error())
		}
	}
	f, ferr := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer f.Close()
	if ferr != nil {
		panic(ferr.Error())
	}
	for i := 0; i < b.N; i++ {
		byteWrite(f, str)
		f.Close()
		f, ferr = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if ferr != nil {
			panic(ferr.Error())
		}
	}
}

func BenchmarkByteWrite(b *testing.B) {
	str := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	fileName := "tmp.log"
	ferr := os.Remove(fileName)
	if ferr != nil {
		if !os.IsNotExist(ferr) {
			panic(ferr.Error())
		}
	}
	f, ferr := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer f.Close()
	if ferr != nil {
		panic(ferr.Error())
	}
	for i := 0; i < b.N; i++ {
		byteWrite(f, str)
	}
}

func BenchmarkStringWrite(b *testing.B) {
	str := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	fileName := "tmp.log"
	ferr := os.Remove(fileName)
	if ferr != nil {
		if !os.IsNotExist(ferr) {
			panic(ferr.Error())
		}
	}
	f, ferr := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer f.Close()
	if ferr != nil {
		panic(ferr.Error())
	}
	for i := 0; i < b.N; i++ {
		stringWrite(f, str)
	}
}

func BenchmarkBufioWrite(b *testing.B) {
	str := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	fileName := "tmp.log"
	ferr := os.Remove(fileName)
	if ferr != nil {
		if !os.IsNotExist(ferr) {
			panic(ferr.Error())
		}
	}
	f, ferr := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer f.Close()
	if ferr != nil {
		panic(ferr.Error())
	}
	w := bufio.NewWriterSize(f, 8192)
	for i := 0; i < b.N; i++ {
		bufioWrite(w, f, str)
	}
}

func BenchmarkDirectWrite(b *testing.B) {
	str := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	fileName := "tmp.log"
	ferr := os.Remove(fileName)
	if ferr != nil {
		if !os.IsNotExist(ferr) {
			panic(ferr.Error())
		}
	}
	f, ferr := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer f.Close()
	if ferr != nil {
		panic(ferr.Error())
	}
	for i := 0; i < b.N; i++ {
		directWrite(f, str)
	}
}
