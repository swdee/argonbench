package main

import "testing"

func BenchmarkAHash1024(b *testing.B) {

	b.ReportAllocs()

	pass := RandStringRunes(1024)
	salt := []byte("sdfggertrtwgrthe")

	for i := 0; i < b.N; i++ {
		_ = AHash(pass, salt)
	}
}

func BenchmarkAHash25(b *testing.B) {

	b.ReportAllocs()

	pass := RandStringRunes(25)
	salt := []byte("sdfggertrtwgrthe")

	for i := 0; i < b.N; i++ {
		_ = AHash(pass, salt)
	}
}
