all: test

test:
	go test ./mm/...
build:
	go build -o server ./example/...

run: build
	./server

