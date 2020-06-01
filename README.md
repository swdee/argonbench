# Quick BenchMark of Argon2 Hashing

## Run Benchmark

```
go test -bench=. -benchtime=10s
```

Output compares hashing of password length of 1024 characters versus 25 characters.

```
BenchmarkAHash1024-8         316          37207613 ns/op        67116783 B/op         33 allocs/op
BenchmarkAHash25-8           315          39111558 ns/op        67115769 B/op         33 allocs/op
```
