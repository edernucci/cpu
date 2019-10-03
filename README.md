# cpu
6502 implementation in Go

```
Running tool: /usr/bin/go test -benchmem -run=^$ github.com/edernucci/cpu -bench . -coverprofile=/tmp/vscode-goFEFJRA/go-code-cover

goos: linux
goarch: amd64
pkg: github.com/edernucci/cpu
BenchmarkLDA-8   	2000000000	         0.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkLDX-8   	2000000000	         0.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkLDY-8   	2000000000	         0.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkSTA-8   	2000000000	         0.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkTXA-8   	2000000000	         0.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkINX-8   	2000000000	         1.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkINY-8   	2000000000	         1.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkTAX-8   	2000000000	         0.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkTYA-8   	2000000000	         0.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkTAY-8   	2000000000	         0.55 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 68.8% of statements
ok  	github.com/edernucci/cpu	15.786s
Success: Benchmarks passed.
```
