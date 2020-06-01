package main

import (
	"fmt"
	"golang.org/x/crypto/argon2"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	pass := RandStringRunes(1024)
	salt := []byte("sdfggertrtwgrthe")

	fmt.Println("pass=", pass)

	start := time.Now()

	PrintMemUsage()
	pw := AHash(pass, salt)
	PrintMemUsage()

	diff := time.Since(start).String()
	fmt.Println("time=", diff, "  hash=", pw)

}

// Hash em
func AHash(pass string, salt []byte) []byte {
	return argon2.IDKey([]byte(pass), salt, 1, 64*1024, 4, 32)
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
