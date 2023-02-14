go: *.go encoding.go
	go test -cover .
	go install

encoding.go: encoding.rl
	ragel -Z -G1 encoding.rl

bench:
	go test -bench=. -count=5  -benchmem

graph:
	ragel -Vp encoding.rl -o encoding.dot
	dot -Tsvg encoding.dot >| encoding.svg

.PHONY: bench graph
