build:
	go build -o ghozz ./cmd

clean:
	rm ghozz

test-all:
	go test -v -count=1 ./...

test-fuzzer:
	go test -v -count=1 ./internal/fuzzer/

test-output:
	go test -v -count=1 ./pkg/output/
