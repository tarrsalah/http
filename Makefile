all: test

test:
	go test ./...

build:
	go build ./...
	go install ./...

example:
	go build -o server ./example/...

run: example
	./server

.PHONY: example
