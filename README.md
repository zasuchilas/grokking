# Grokking


---

### Benchmarks

```shell

go test -bench=. ./...

go install golang.org/x/perf/cmd/benchstat@latest
go test -bench=. ./search/simple/ -count=10 >./search/simple_search.txt
go test -bench=. ./search/binary/ -count=10 >./search/binary_search.txt
benchstat ./search/simple_search.txt ./search/binary_search.txt

```
