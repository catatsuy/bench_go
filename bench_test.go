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

func BenchmarkInit文字列をバイト列に変換(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
}

func BenchmarkInitバイト列を初期化(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}
	}
}

func BenchmarkInitバイト列を3回に分けて初期化(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = append(append(append(make([]byte, 0, 200), []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}...), []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}...), []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}...)
	}
}

func BenchmarkInitバイト列を100回に分けて初期化(b *testing.B) {
	b.ReportAllocs()
	var test_b []byte
	for i := 0; i < b.N; i++ {
		test_b = make([]byte, 0, 200)
		for i := 0; i < 100; i++ {
			test_b = append(test_b, 'a')
		}
	}
}

func BenchmarkAppendStringByteAlloc(b *testing.B) {
	test_b := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	for i := 0; i < b.N; i++ {
		m2 := make([]byte, 0, 200)
		m2 = append(m2, test_b...)
	}
}

func BenchmarkAppendStringByteDByte(b *testing.B) {
	test_b := []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}
	for i := 0; i < b.N; i++ {
		m2 := make([]byte, 0, 200)
		m2 = append(m2, test_b...)
	}
}

func BenchmarkAppendStringByteZero(b *testing.B) {
	test_b := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	for i := 0; i < b.N; i++ {
		m2 := make([]byte, 0)
		m2 = append(m2, test_b...)
	}
}
