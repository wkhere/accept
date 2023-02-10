go:
	go vet
	go test -cover .
	go install

bench:
	go test -bench=. -count=5  -benchmem

.PHONY: go bench
